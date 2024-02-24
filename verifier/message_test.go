package verifier

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type SccSigTestSuite struct {
	suite.Suite

	b *backend.TestBackend

	sccApproveMsg *Message
	sccRejectMsg  *Message

	l2ooApproveMsg *Message
	l2ooRejectMsg  *Message
}

func TestSccSig(t *testing.T) {
	suite.Run(t, new(SccSigTestSuite))
}

func (s *SccSigTestSuite) SetupSuite() {
	s.b = backend.NewTestBackend()

	s.sccApproveMsg = NewSccMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		true,
	)
	s.sccRejectMsg = NewSccMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		false,
	)

	s.l2ooApproveMsg = NewL2ooMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		big.NewInt(3),
		big.NewInt(4),
		true,
	)
	s.l2ooRejectMsg = NewL2ooMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		big.NewInt(3),
		big.NewInt(4),
		false,
	)
}

func (s *SccSigTestSuite) TestNewApproveSccMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
		"01",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.sccApproveMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.sccApproveMsg.Eip712Msg)
}

func (s *SccSigTestSuite) TestNewApproveL2ooMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		strings.Join([]string{
			"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
			"00000000000000000000000000000003",
			"00000000000000000000000000000004",
		}, ""),
		"01",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.l2ooApproveMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.l2ooApproveMsg.Eip712Msg)
}

func (s *SccSigTestSuite) TestNewRejectSccMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
		"00",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.sccRejectMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.sccRejectMsg.Eip712Msg)
}

func (s *SccSigTestSuite) TestNewRejectL2ooMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		strings.Join([]string{
			"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
			"00000000000000000000000000000003",
			"00000000000000000000000000000004",
		}, ""),
		"00",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.l2ooRejectMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.l2ooRejectMsg.Eip712Msg)
}

func (s *SccSigTestSuite) TestSignature() {
	got1, _ := s.sccApproveMsg.Signature(s.b.SignData)
	got2, _ := s.sccRejectMsg.Signature(s.b.SignData)
	got3, _ := s.l2ooApproveMsg.Signature(s.b.SignData)
	got4, _ := s.l2ooRejectMsg.Signature(s.b.SignData)

	s.Equal(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), got1[:])
	s.Equal(hexutil.MustDecode(
		"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
			"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), got2[:])
	s.Equal(hexutil.MustDecode(
		"0x99fc461006773b92267fb7c52f55322e3bdecbf3120d76ffec252c9b8381bba6"+
			"6397084757d9027ba693f2b5eea4766edd3a2ccedf29c55bfed51d924a260b701b"), got3[:])
	s.Equal(hexutil.MustDecode(
		"0xca30829bdaf7e398d239ee911b8ee307b3f5d3e4f0f0bf0372806d0778c00a28"+
			"1f986ef86fd2ff93b50591cebd659a6d6b18a921e7a6e6cd68c91a30a03a470b1b"), got4[:])
}

func (s *SccSigTestSuite) TestEcrecover() {
	got1, _ := s.sccApproveMsg.Ecrecover(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0" +
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"))
	got2, _ := s.sccRejectMsg.Ecrecover(hexutil.MustDecode(
		"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729" +
			"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"))
	got3, _ := s.l2ooApproveMsg.Ecrecover(hexutil.MustDecode(
		"0x99fc461006773b92267fb7c52f55322e3bdecbf3120d76ffec252c9b8381bba6" +
			"6397084757d9027ba693f2b5eea4766edd3a2ccedf29c55bfed51d924a260b701b"))
	got4, _ := s.l2ooRejectMsg.Ecrecover(hexutil.MustDecode(
		"0xca30829bdaf7e398d239ee911b8ee307b3f5d3e4f0f0bf0372806d0778c00a28" +
			"1f986ef86fd2ff93b50591cebd659a6d6b18a921e7a6e6cd68c91a30a03a470b1b"))
	got5, _ := s.sccRejectMsg.Ecrecover(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729" +
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef10"))

	s.Equal(s.b.Signer(), got1)
	s.Equal(s.b.Signer(), got2)
	s.Equal(s.b.Signer(), got3)
	s.Equal(s.b.Signer(), got4)
	s.NotEqual(s.b.Signer(), got5)
}

func (s *SccSigTestSuite) TestVerifySigner() {
	got1 := s.sccApproveMsg.VerifySigner(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), s.b.Signer())
	got2 := s.sccRejectMsg.VerifySigner(hexutil.MustDecode(
		"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
			"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), s.b.Signer())
	got3 := s.l2ooApproveMsg.VerifySigner(hexutil.MustDecode(
		"0x99fc461006773b92267fb7c52f55322e3bdecbf3120d76ffec252c9b8381bba6"+
			"6397084757d9027ba693f2b5eea4766edd3a2ccedf29c55bfed51d924a260b701b"), s.b.Signer())
	got4 := s.l2ooRejectMsg.VerifySigner(hexutil.MustDecode(
		"0xca30829bdaf7e398d239ee911b8ee307b3f5d3e4f0f0bf0372806d0778c00a28"+
			"1f986ef86fd2ff93b50591cebd659a6d6b18a921e7a6e6cd68c91a30a03a470b1b"), s.b.Signer())
	got5 := s.sccApproveMsg.VerifySigner(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), common.Address{})

	s.Nil(got1)
	s.Nil(got2)
	s.Nil(got3)
	s.Nil(got4)
	s.ErrorContains(got5, "signer mismatch")
}
