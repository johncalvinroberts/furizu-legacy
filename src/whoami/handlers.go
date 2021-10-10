package whoami

import (
	"fmt"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
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
	token, err := UpsertWhoamiChallenge(req.Email)

	if err != nil {
		fmt.Printf("token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Printf("token: %s", token)
	// TODO: send email {answer}
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
	result, err := FindWhoamiChallenge(req.Token)
	if err != nil && strings.Contains(err.Error(), "no item found") {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Token does not exist"})
		return
	}

	if err != nil {
		fmt.Printf("err %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// return 400 if invalid
	if result.Exp.Before(time.Now()) || result.Email != req.Email {
		c.JSON(http.StatusBadRequest,
			map[string]interface{}{"success": false, "message": "Token invalid or expired"})
		return
	}
	/* TODO:
	delete token from dynamo
	create/update user
	issue JWT

	*/
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Refresh(c *gin.Context) {
	// issue a new token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func Revoke(c *gin.Context) {
	// destroy token
	c.JSON(http.StatusOK, map[string]bool{"success": true})
}
