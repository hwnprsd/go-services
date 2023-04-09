package controllers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"flaq.club/api/models"
	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

type Claims struct {
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func (claims *Claims) Valid() error {
	// Check if valid
	return nil
}

func (c *Controller) SetupAuthRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/auth")
	goth.UseProviders(
		google.New(os.Getenv("GAUTH_CLIENT_KEY"), os.Getenv("GAUTH_SECRET"), os.Getenv("GAUTH_CALLBACK_URL")),
	)
	jwksURL := "https://www.googleapis.com/oauth2/v3/certs"
	c.jwk, _ = keyfunc.Get(jwksURL, keyfunc.Options{})
	fiberw.RawGet(group, "/login/:provider", c.AuthLogin).WithParam("provider")
	fiberw.RawPost(group, "/logout", c.AuthLogout)
	fiberw.RawPost(group, "/callback/:provider", c.AuthCallback).WithParam("provider")
	// fiberw.RawGet(group, "/check", c.CheckUser)
}

func (c *Controller) InjectUser(ctx *fiber.Ctx) (*models.User, error) {
	val := ctx.Cookies("googleauth")
	log.Println(val)
	if val == "" {
		return nil, fiberw.NewRequestError(402, "Not Authorized to Access this route", errors.New("No auth token found"))
	}
	claims := Claims{}
	_, err := jwt.ParseWithClaims(val, &claims, c.jwk.Keyfunc)
	if err != nil {
		return nil, fiberw.NewRequestError(402, "Not Authorized to Access this route", err)
	}
	user, err := c.getUser(claims.Sub)
	if err != nil {
		return nil, fiberw.NewRequestError(402, "Not Authorized to Access this route", err)
	}
	return user, nil
}

func (c *Controller) AuthLogin(ctx *fiber.Ctx) error {
	err := goth_fiber.BeginAuthHandler(ctx)
	log.Println(err)
	return ctx.SendString("Idk")
}

func (c *Controller) AuthCallback(ctx *fiber.Ctx) error {
	log.Println("Handling callback?")
	user, err := goth_fiber.CompleteUserAuth(ctx)
	now := time.Now()
	user.ExpiresAt = now.Add(24 * 30 * time.Hour)
	ctx.Cookie(&fiber.Cookie{
		Name:     "googleauth",
		HTTPOnly: true,
		Value:    user.IDToken,
		Expires:  user.ExpiresAt,
	})
	log.Println(user.ExpiresAt)
	log.Println(time.Now())
	claims := Claims{}
	_, err = jwt.ParseWithClaims(user.IDToken, &claims, c.jwk.Keyfunc)
	userData := models.User{
		Name:              user.Name,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             &user.Email,
		UserID:            user.UserID,
		UniqueId:          claims.Sub,
		AvatarURL:         user.AvatarURL,
		NickName:          user.NickName,
		Location:          user.Location,
		AccessTokenSecret: user.AccessTokenSecret,
		RefreshToken:      user.RefreshToken,
		ExpiresAt:         user.ExpiresAt,
	}
	u, err := c.GetOrCreateUser(&userData)
	if err != nil {
		return err
	}
	return ctx.SendString(fmt.Sprintf("User exists with Email %s", *u.Email))
}

func (c *Controller) AuthLogout(ctx *fiber.Ctx) error {
	if err := goth_fiber.Logout(ctx); err != nil {
		log.Fatal(err)
	}
	return ctx.SendString("Successfully Loggedout")
}
