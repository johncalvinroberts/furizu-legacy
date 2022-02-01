package whoami

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/furizu/app/users"
	"github.com/johncalvinroberts/furizu/app/utils"
)

type TokenSet struct {
	AccessToken  string
	RefreshToken string
}

// get current user
func Me(ctx context.Context) (user *users.User, err error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	userPartial, err := utils.Authenticate(gc)
	if err != nil {
		return nil, err
	}
	user, err = users.FindUserByEmail(userPartial.Email)
	if err != nil {
		return nil, err
	}
	return user, err
}

// initialize a whoami flow
func Start(email string) error {
	// save to db
	token, err := upsertWhoamiChallenge(email)
	log.Print(token)
	if err != nil {
		log.Printf("token: %s", err.Error())
		return err
	}
	msg := fmt.Sprintf(`Copy and paste this temporary login code
	<pre style="padding:16px 24px;border:1px solid #eeeeee;background-color:#f4f4f4;border-radius:3px;font-family:monospace;margin-bottom:24px">%s</pre>
	`, token)
	utils.SendANiceEmail(email, msg, "Log in code to furizu")
	return nil
}

// find WhoamiChallenge based on request
func Redeem(email string, token string) (*TokenSet, error) {
	// lookup whoami challenge
	result, err := findWhoamiChallenge(token)
	if err != nil && strings.Contains(err.Error(), "no item found") {
		return nil, errors.New("token invalid or expired")
	}

	if err != nil {
		log.Printf("err %s", err.Error())
		return nil, err
	}
	// check if whoamireq is expired
	if result.Exp.Before(time.Now()) || result.Email != email {
		return nil, errors.New("token invalid or expired")
	}
	// if we get here, successfully redeemed token
	// create/update user
	user, err := users.UpsertUser(email)
	if err != nil {
		log.Printf("Failed to upsert user %v", err)
		return nil, err
	}
	// create tokens
	accessToken, err := utils.FurizuJWT.GenerateAccessToken(map[string]string{
		"userId": fmt.Sprint(user.ID),
		"email":  fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue accessToken %v", err)
		return nil, err
	}
	// TODO: persist session in db?
	refreshToken, err := utils.FurizuJWT.GenerateRefreshToken(map[string]string{
		"userId": fmt.Sprint(user.ID),
		"email":  fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue refreshToken %v", err)
		return nil, err
	}
	// lastly, delete token from dynamo
	err = destroyWhoamiChallenge(token)
	if err != nil {
		log.Printf("Failed to destroy token %v", err)
		return nil, err
	}
	return &TokenSet{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func Refresh(prevRefreshToken string) (*TokenSet, error) {
	decoded, err := utils.FurizuJWT.ValidateFromToken(prevRefreshToken)
	if err != nil {
		return nil, err
	}
	email, err := decoded.Get("email")
	if err != nil {
		return nil, err
	}
	// TODO: check session in db
	user, err := users.FindUserByEmail(email.(string))
	if err != nil {
		log.Printf("Failed to find user during refresh token %v", err)
		return nil, err
	}
	// create tokens
	accessToken, err := utils.FurizuJWT.GenerateAccessToken(map[string]string{
		"userId": fmt.Sprint(user.ID),
		"email":  fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue accessToken %v", err)
		return nil, err
	}
	refreshToken, err := utils.FurizuJWT.GenerateRefreshToken(map[string]string{
		"userId": fmt.Sprint(user.ID),
		"email":  fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue refresh token %v", err)
	}
	return &TokenSet{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func Revoke(c *gin.Context) {
	// destroy token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}
