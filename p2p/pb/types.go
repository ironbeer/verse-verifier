package pb

type ISignature interface {
	GetId() string         // signature id
	GetPreviousId() string // previous signature id
	GetSigner() []byte     // signer address
}

type ISignatureRequest interface {
	GetIdAfter() string // signature id
	GetSigner() []byte  // signer address
}

type ICommonSignatureRequest interface {
	GetId() string         // signature id
	GetPreviousId() string // previous signature id
}
