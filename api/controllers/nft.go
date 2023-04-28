package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
	"gorm.io/gorm/clause"
)

func (c *Controller) SetupNFTRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/nft")

	fiberw.Post(group, "/poap/mint", c.MintPOAP)
	fiberw.Post(group, "/scw/create", c.CreateSCW)
	fiberw.Post(group, "/scw/relay", c.RelayTx)
}

type MintNftBody struct {
	Email         string `json:"email" validate:"required,email,min=6,max=32"`
	Name          string `json:"name" validate:"required,min=3,max=23"`
	WalletAddress string `validate:"required,min=42,max=42" json:"wallet_address"`
	MintSecret    string `json:"mint_secret" validate:"required"`
	EventId       uint   `json:"event_id" validate:"required"`
}

func (ctrl *Controller) MintPOAP(body MintNftBody) (*models.Task, error) {
	secret := os.Getenv("POAP_MINT_SECRET")
	if body.MintSecret != secret {
		return nil, &fiberw.RequestError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid Secret",
			Err:        errors.New("Invalid secret provided for minting"),
		}
	}

	existingUser := models.User{}

	dbRes := ctrl.DB.Clauses(clause.Returning{}).Where("email = ?", body.Email).First(&existingUser)
	if dbRes.Error != nil {
		existingUser.Email = &body.Email
		existingUser.Level = 1
		dbRes = ctrl.DB.Create(&existingUser)

		if dbRes.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: 400,
				Message:    "Error creating / finding user with the email address",
				Err:        dbRes.Error,
			}
		}
	}

	job := models.Task{
		UserID:   existingUser.ID,
		Status:   status.POAP_REQUESTED,
		Category: "POAP",
	}
	ctrl.DB.Clauses(clause.Returning{}).Create(&job)

	log.Println("New Job", job.ID)
	// payload := shared_types.NewMintPoapMessage(body.Email, body.WalletAddress, body.Name, body.TokenURI, 1)
	payload2 := shared_types.NewCreateGifMessage(job.ID, 1, body.EventId, body.Name, body.WalletAddress, body.Email)
	ctrl.MQ.GifQueue.PublishMessage(*payload2)
	// ctrl.MQ.NftQueue.PublishMessage(*payload)
	return &job, nil
}

type CreateSCWBody struct {
	OwnerAddress string `json:"owner_address"`
}

func (ctrl *Controller) CreateSCW(body CreateSCWBody) (*string, error) {
	done := "done"
	payload := shared_types.CreateSmartContractWalletMessage(1, body.OwnerAddress)
	ctrl.MQ.NftQueue.PublishMessage(payload)
	return &done, nil
}

type RelayTxBody struct {
	UserAddress     string `json:"user_address"`
	ContractAddress string `json:"contract_address"`
	Data            string `json:"data"`
	Signature       string `json:"signature"`
	Nonce           int64  `json:"nonce"`
}

func (ctrl *Controller) RelayTx(body RelayTxBody) (*string, error) {
	done := "done"
	log.Println("Data", body.Data)
	log.Println("Sign", body.Signature)
	log.Println("Nonce", body.Nonce)
	payload := shared_types.CreateRelayTxMessage(1, body.UserAddress, body.ContractAddress, body.Data, body.Signature, body.Nonce)
	ctrl.MQ.NftQueue.PublishMessage(payload)
	return &done, nil
}
