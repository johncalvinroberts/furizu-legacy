package whoami

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/furizu/src/users"
	"github.com/johncalvinroberts/furizu/src/utils"
)

type StartWhoamiReq struct {
	Email string `json:"email"`
}

type RedeemWhoamiReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// get current user
func Me(c *gin.Context) {
	// lookup user
	c.JSON(http.StatusAccepted, map[string]bool{"success": true})
}

// initialize a whoami flow
func Start(c *gin.Context) {
	// lookup existing user
	req := &StartWhoamiReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// save to db
	token, err := upsertWhoamiChallenge(req.Email)
	log.Print(token)
	if err != nil {
		log.Printf("token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	msg := fmt.Sprintf(`Copy and paste this temporary login code
	<pre style="padding:16px 24px;border:1px solid #eeeeee;background-color:#f4f4f4;border-radius:3px;font-family:monospace;margin-bottom:24px">%s</pre>
	`, token)
	utils.SendANiceEmail(req.Email, msg, "Log in code to furizu")
	c.JSON(http.StatusAccepted, map[string]bool{"success": true})
}

// find WhoamiChallenge based on request
func Redeem(c *gin.Context) {
	req := &RedeemWhoamiReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// lookup whoami challenge
	result, err := findWhoamiChallenge(req.Token)
	if err != nil && strings.Contains(err.Error(), "no item found") {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Token does not exist"})
		return
	}

	if err != nil {
		log.Printf("err %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// return 400 if invalid
	if result.Exp.Before(time.Now()) || result.Email != req.Email {
		c.JSON(http.StatusBadRequest,
			map[string]interface{}{"success": false, "message": "Token invalid or expired"})
		return
	}
	// if we get here, successfully redeemed token
	// create/update user
	user, err := users.UpsertUser(req.Email)
	if err != nil {
		log.Printf("Failed to upsert user %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// issue jwt
	token, err := utils.FurizuJWT.ToToken(map[string]string{
		"id":    fmt.Sprint(user.Id),
		"email": fmt.Sprint(user.Email),
	})
	if err != nil {
		log.Printf("Failed to issue jwt %v", err)
	}
	// lastly, delete token from dynamo
	err = destroyWhoamiChallenge(req.Token)
	if err != nil {
		log.Printf("Failed to destroy token %v", err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{"success": true, "token": token})
}

func Refresh(c *gin.Context) {
	// issue a new token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Revoke(c *gin.Context) {
	// destroy token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}
