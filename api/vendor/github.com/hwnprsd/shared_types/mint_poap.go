package shared_types

type MintPoapMessage struct {
	MessagingBase
	Email    string
	Address  string
	Name     string
	TokenURI string
}

func NewMintPoapMessage(email string, walletAddress string, name string, tokenUri string) *MintPoapMessage {
	return &MintPoapMessage{
		MessagingBase: MessagingBase{
			WorkType: "MINT_POAP",
		},
		Name:     name,
		Email:    email,
		Address:  walletAddress,
		TokenURI: tokenUri,
	}
}
