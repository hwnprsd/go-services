package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hwnprsd/shared_types"
	"gorm.io/gorm/clause"
)

func (c Controller) SetupQuizRoutes() {
	group := c.FiberApp.Group("quiz")

	group.Post("/submit", func(ctx *fiber.Ctx) error {
		body := new(SubmitQuizParticipationBody)
		return utils.PostRequestHandler(ctx, c.SubmitQuizParticipation(), utils.RequestBody{Data: body})
	})

	group.Post("/request-email", func(ctx *fiber.Ctx) error {
		body := new(NFTClaimEmailBody)
		return utils.PostRequestHandler(ctx, c.RequestNFTClaimEmail(), utils.RequestBody{Data: body})
	})

	group.Get("/claim-info", func(ctx *fiber.Ctx) error {
		quizClaimId := ctx.Query("quizClaimId")
		query := new(GetSubmissionInfoQuery)
		query.QuizClaimID = quizClaimId
		return utils.GetRequestHandler(ctx, c.GetSubmissionInfo(), utils.RequestBody{Data: query})
	})

	group.Post("/mint", func(ctx *fiber.Ctx) error {
		body := new(NFTClaimAttempt)
		return utils.PostRequestHandler(ctx, c.MintQuizNFT(), utils.RequestBody{Data: body})
	})
}

type SubmitQuizParticipationBody struct {
	Email           string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	QuizId          string `json:"quiz_id,omitempty" validate:"required"`
	ShouldSendEmail bool   `json:"should_send_email"`
}

func (ctrl *Controller) SubmitQuizParticipation() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*SubmitQuizParticipationBody)
		tokenId := uuid.New().String()
		quizSubmission := models.QuizSubmission{
			Email:                body.Email,
			QuizID:               body.QuizId,
			ClaimID:              tokenId,
			IsNFTClaimed:         false,
			IsNFTClaimMailSent:   body.ShouldSendEmail,
			NFTClaimAttemptCount: 0,
		}

		result := ctrl.DB.Clauses(clause.Returning{}).Create(&quizSubmission)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Creating Quiz Submission",
				Err:        result.Error,
			}
		}

		if body.ShouldSendEmail {
			mailerMessage := shared_types.NewSendMailMessage(
				0,
				quizSubmission.Email,
				"Congratulations! You have received a completion Insignia NFT from Flaq",
				2,
				map[string]string{
					"claim_url":        fmt.Sprintf("https://learn.flaq.club/claim-nft/%s", tokenId),
					"claim_url_mobile": fmt.Sprintf("https://metamask.app.link/dapp/learn.flaq.club/claim-nft/%s", tokenId),
				},
			)

			log.Println("Sending Claim Email")
			// Handle errors
			err := ctrl.MQ.MailerQueue.PublishMessage(mailerMessage)
			if err != nil {
				log.Println("Error sending mail")
			}
		}

		return quizSubmission, nil
	}
}

type NFTClaimEmailBody struct {
	QuizClaimID string `json:"quiz_claim_id,omitempty" validate:"required"`
}

func (ctrl *Controller) RequestNFTClaimEmail() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*NFTClaimEmailBody)
		result := ctrl.DB.Model(&models.QuizSubmission{}).Where("claim_id = ? AND is_nft_claim_mail_sent = ?", body.QuizClaimID, false).Update("is_nft_claim_mail_sent", true)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        result.Error,
				Message:    "Error updating the entry",
			}
		}

		if result.RowsAffected == 0 {
			return nil, &utils.RequestError{
				StatusCode: 400,
				Err:        errors.New("Invalid claim ID"),
				Message:    "The claim ID is invalid or an email has already sent",
			}
		}

		quizSubmission := models.QuizSubmission{}
		ctrl.DB.Model(&quizSubmission).Where("claim_id = ?", body.QuizClaimID).First(&quizSubmission)

		mailerMessage := shared_types.NewSendMailMessage(
			0,
			quizSubmission.Email,
			"Your hard earned NFT's are here",
			2,
			map[string]string{
				"claim_url": body.QuizClaimID,
			},
		)

		// Handle errors
		err := ctrl.MQ.MailerQueue.PublishMessage(mailerMessage)
		if err != nil {
			log.Println("Error sending mail")
		}
		return "Claim Mail Sent", nil
	}
}

type GetSubmissionInfoQuery struct {
	QuizClaimID string `json:"quiz_claim_id,omitempty" validate:"required"`
}

func (ctrl *Controller) GetSubmissionInfo() utils.GetHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		query := data.Data.(*GetSubmissionInfoQuery)
		submission := models.QuizSubmission{}
		result := ctrl.DB.Model(&submission).Where("claim_id = ?", query.QuizClaimID).First(&submission)

		// Consider using transactions to make this faster
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error fetching data",
				Err:        result.Error,
			}
		}

		result = ctrl.DB.Model(&models.QuizSubmission{}).Where("claim_id = ?", query.QuizClaimID).Update("NFTClaimAttemptCount", submission.NFTClaimAttemptCount+1)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        result.Error,
				Message:    "Error updating the entry",
			}
		}

		return submission, nil
	}
}

type NFTClaimAttempt struct {
	QuizClaimID   string `json:"quiz_claim_id,omitempty" validate:"required"`
	WalletAddress string `json:"wallet_address,omitempty" validate:"required"`
}

func (ctrl *Controller) MintQuizNFT() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*NFTClaimAttempt)
		submission := models.QuizSubmission{}
		result := ctrl.DB.Model(&submission).Where("claim_id = ?", body.QuizClaimID).First(&submission)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error fetching data",
				Err:        result.Error,
			}
		}

		if submission.IsNFTClaimed {
			return nil, &utils.RequestError{
				StatusCode: http.StatusBadRequest,
				Message:    "NFT is already claimed for this Claim ID",
				Err:        errors.New("Error claiming NFT"),
			}
		}

		// TODO: Create or get NFT Token URI
		tokenUri := "https://bafybeibb7ddws3smwceufpdgkj4zyqmirwfzgtbfubwrsp4gmebjkxbtde.ipfs.nftstorage.link/diveIntoWeb3_track.json"
		// TOOD: Update
		message := shared_types.NewMintQuizNFTMessage(0, submission.Email, body.WalletAddress, tokenUri)
		if err := ctrl.MQ.NftQueue.PublishMessage(message); err != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error sending message across services",
				Err:        err,
			}
		}

		// Expecting no errors at this point
		ctrl.DB.Model(&models.QuizSubmission{}).Where("claim_id = ?", submission.ClaimID).Update("is_nft_claimed", true)

		return "NFT Minting Started - Email will be sent upon completion", nil

	}
}
