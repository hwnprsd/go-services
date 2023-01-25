package shared_types

const WORK_TYPE_MINT_POAP = "MINT_POAP"

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
			WorkType: WORK_TYPE_MINT_POAP,
		},
		Name:     name,
		Email:    email,
		Address:  walletAddress,
		TokenURI: tokenUri,
	}
}

const WORK_TYPE_MINT_QUIZ_NFT = "MINT_QUIZ_NFT"

type MintQuizNFTMessage struct {
	MessagingBase
	Address  string
	TokenURI string
	Email    string
}

func NewMintQuizNFTMessage(email string, walletAddress string, tokenUri string) *MintQuizNFTMessage {
	return &MintQuizNFTMessage{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_MINT_QUIZ_NFT,
		},
		Email:    email,
		Address:  walletAddress,
		TokenURI: tokenUri,
	}
}
