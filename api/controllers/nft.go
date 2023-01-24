package controllers

import (
	"errors"
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/google/uuid"
	"github.com/hwnprsd/shared_types"
)

type MintNftBody struct {
	Email         string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	Name          string `json:"name,omitempty" validate:"required,min=3,max=12"`
	WalletAddress string `json:"walletAddress,omitempty" validate:"required,min=42,max=42"`
}

func (ctrl *Controller) MintPOAP() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*MintNftBody)
		payload := shared_types.NewMintPoapMessage(body.Email, body.WalletAddress, body.Name, "https://google.com")
		ctrl.MQ.NftQueue.PublishMessage(*payload)
		return "NFT Minting Started. Check back on your email", nil
	}
}

type SubmitQuizParticipationBody struct {
	Email  string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	QuizId string `json:"quiz_id,omitempty" validate:"required"`
}

func (ctrl *Controller) SubmitQuizParticipation() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*SubmitQuizParticipationBody)
		tokenId := uuid.New().String()
		quizSubmission := models.QuizSubmission{
			Email:                &body.Email,
			QuizID:               body.QuizId,
			ClaimID:              tokenId,
			IsNFTClaimed:         false,
			IsNFTClaimMailSent:   false,
			NFTClaimAttemptCount: 0,
		}
		result := ctrl.DB.Create(&quizSubmission)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Creating Quiz Submission",
				Err:        result.Error,
			}
		}

		return tokenId, nil
	}
}

type RequestNFTClaimEmailBody struct {
	QuizClaimID string `json:"quiz_claim_id,omitempty" validate:"required"`
}

func (ctrl *Controller) RequestNFTClaimEmail() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*RequestNFTClaimEmailBody)
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

		ctrl.MQ.MailerQueue.PublishMessage(utils.Map{
			"type":    "Send Mail",
			"content": "Yay",
		})
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

		if result.Error != nil {

			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error fetching data",
				Err:        result.Error,
			}
		}

		return submission, nil
	}
}
