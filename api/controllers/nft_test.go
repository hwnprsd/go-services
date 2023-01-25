package controllers

import (
	"reflect"
	"testing"

	"flaq.club/api/app"
	"flaq.club/api/utils"
)

func TestController_MintPOAP(t *testing.T) {
	type fields struct {
		App *app.App
	}
	tests := []struct {
		name   string
		fields fields
		want   utils.PostHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &Controller{
				App: tt.fields.App,
			}
			if got := ctrl.MintPOAP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.MintPOAP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_SubmitQuizParticipation(t *testing.T) {
	type fields struct {
		App *app.App
	}
	tests := []struct {
		name   string
		fields fields
		want   utils.PostHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &Controller{
				App: tt.fields.App,
			}
			if got := ctrl.SubmitQuizParticipation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.SubmitQuizParticipation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_RequestNFTClaimEmail(t *testing.T) {
	type fields struct {
		App *app.App
	}
	tests := []struct {
		name   string
		fields fields
		want   utils.PostHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &Controller{
				App: tt.fields.App,
			}
			if got := ctrl.RequestNFTClaimEmail(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.RequestNFTClaimEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_GetSubmissionInfo(t *testing.T) {
	type fields struct {
		App *app.App
	}
	tests := []struct {
		name   string
		fields fields
		want   utils.GetHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &Controller{
				App: tt.fields.App,
			}
			if got := ctrl.GetSubmissionInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.GetSubmissionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_MintQuizNFT(t *testing.T) {
	type fields struct {
		App *app.App
	}
	tests := []struct {
		name   string
		fields fields
		want   utils.PostHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &Controller{
				App: tt.fields.App,
			}
			if got := ctrl.MintQuizNFT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.MintQuizNFT() = %v, want %v", got, tt.want)
			}
		})
	}
}
