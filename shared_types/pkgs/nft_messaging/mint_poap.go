package nftmessaging

import "flaq.club/shared_types/pkgs/messaging_base"

type MintPoapMessage struct {
	messaging_base.MessagingBase
	Email   string
	Address string
	Name    string
}

func NewMintPoapMessage(email string, walletAddress string, name string) *MintPoapMessage {
	return &MintPoapMessage{
		MessagingBase: messaging_base.MessagingBase{
			WorkType: "MINT_POAP",
		},
		Name:    name,
		Email:   email,
		Address: walletAddress,
	}
}
