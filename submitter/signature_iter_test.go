package submitter

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
)

func (s *SubmitterTestSuite) TestSignatureIterator() {
	var signers [20]common.Address
	for i := range signers {
		signers[i] = s.RandAddress()
	}

	type signerGroup struct {
		signers []common.Address
		stake   int64
	}
	signerGroups := [4]*signerGroup{
		{signers: signers[0:5], stake: 5},
		{signers: signers[5:10], stake: 15},
		{signers: signers[15:20], stake: 50}, // most stakes
		{signers: signers[10:15], stake: 20},
	}

	type signatureGroup struct {
		sg         *signerGroup
		rollupHash common.Hash
		approved   bool
		signatures []*database.OptimismSignature
	}
	rollupHash0 := s.RandHash()
	rollupHash1 := s.RandHash()
	rollupHash2 := s.RandHash()
	rollupHash3 := s.RandHash()
	cases := [3][]*signatureGroup{
		// rollupIndex = 0
		{
			{sg: signerGroups[1], rollupHash: rollupHash0, approved: true},
			{sg: signerGroups[2], rollupHash: rollupHash0, approved: false}, // want
			{sg: signerGroups[3], rollupHash: rollupHash1, approved: true},
			{sg: signerGroups[0], rollupHash: rollupHash1, approved: false},
		},
		// rollupIndex = 1
		{
			{sg: signerGroups[3], rollupHash: rollupHash2, approved: true},
			{sg: signerGroups[0], rollupHash: rollupHash2, approved: false},
			{sg: signerGroups[1], rollupHash: rollupHash3, approved: true},
			{sg: signerGroups[2], rollupHash: rollupHash3, approved: false}, // want
		},
		// rollupIndex = 2 (should return `*StakeAmountShortage`)
		{
			{sg: signerGroups[0], rollupHash: s.RandHash(), approved: true},
		},
	}
	want0 := cases[0][1]
	want1 := cases[1][3]

	// setup stakemanager mock
	smock := &testhelper.StakeManagerMock{}
	smcache := stakemanager.NewCache(smock)
	for _, group := range signerGroups {
		mul := big.NewInt(group.stake / int64(len(group.signers)))
		for _, signer := range group.signers {
			smock.Owners = append(smock.Owners, s.RandAddress())
			smock.Operators = append(smock.Operators, signer)
			smock.Stakes = append(smock.Stakes, new(big.Int).Mul(minStake, mul))
			smock.Candidates = append(smock.Candidates, true)
		}
	}
	smcache.Refresh(context.Background())

	// save signatures
	for rollupIndex, c := range cases {
		for _, group := range c {
			for _, signer := range group.sg.signers {
				sig, _ := s.DB.OPSignature.Save(
					nil, nil,
					signer,
					s.SCCAddr,
					uint64(rollupIndex),
					group.rollupHash,
					group.approved,
					database.RandSignature(),
				)
				group.signatures = append(group.signatures, sig)
			}
		}
	}

	sort.Sort(database.OptimismSignatures(want0.signatures))
	sort.Sort(database.OptimismSignatures(want1.signatures))

	iter := &signatureIterator{
		db:           s.DB,
		stakemanager: smcache,
		contract:     s.SCCAddr,
		rollupIndex:  0,
	}

	// assert
	gots0, err0 := iter.next()
	gots1, err1 := iter.next()

	s.Nil(err0)
	s.Nil(err1)

	s.Len(gots0, len(want0.signatures))
	s.Len(gots1, len(want1.signatures))

	for i, want := range want0.signatures {
		s.Equal(want.Signature, gots0[i].Signature)
	}
	for i, want := range want1.signatures {
		s.Equal(want.Signature, gots1[i].Signature)
	}

	// should return `*StakeAmountShortage`
	for i := range smock.Operators {
		smock.Stakes[i] = minStake
	}
	smcache.Refresh(context.Background())
	_, err := iter.next()
	s.ErrorContains(err, "stake amount shortage")
}
