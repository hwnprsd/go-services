package app

import (
	"testing"

	"flaq.club/api/messaging"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TestApp_HealthCheck(t *testing.T) {
	type fields struct {
		DB       gorm.DB
		MQ       *messaging.Messaging
		FiberApp *fiber.App
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				DB:       tt.fields.DB,
				MQ:       tt.fields.MQ,
				FiberApp: tt.fields.FiberApp,
			}
			if err := a.HealthCheck(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("App.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
