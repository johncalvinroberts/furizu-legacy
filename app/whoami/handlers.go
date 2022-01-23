package whoami

import (
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

type RedeemWhoamiReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// get current user
func Me(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.RespondWithError(c, http.StatusInternalServerError, map[string]interface{}{"success": false})
	}
	c.JSON(http.StatusAccepted, map[string]interface{}{"success": true, "user": user})
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
func Redeem(email string, token string) (jwt string, err error) {
	// lookup whoami challenge
	result, err := findWhoamiChallenge(token)
	if err != nil && strings.Contains(err.Error(), "no item found") {
		return jwt, errors.New("token invalid or expired")
	}

	if err != nil {
		log.Printf("err %s", err.Error())
		return jwt, err
	}
	// return 400 if invalid
	if result.Exp.Before(time.Now()) || result.Email != email {
		return jwt, errors.New("token invalid or expired")
	}
	// if we get here, successfully redeemed token
	// create/update user
	user, err := users.UpsertUser(email)
	if err != nil {
		log.Printf("Failed to upsert user %v", err)
		return jwt, err
	}
	// issue jwt
	jwt, err = utils.FurizuJWT.ToToken(map[string]string{
		"id":    fmt.Sprint(user.Id),
		"email": fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue jwt %v", err)
	}
	// lastly, delete token from dynamo
	err = destroyWhoamiChallenge(token)
	if err != nil {
		log.Printf("Failed to destroy token %v", err)
	}
	return jwt, nil
}

func Refresh(c *gin.Context) {
	// issue a new token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Revoke(c *gin.Context) {
	// destroy token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}
