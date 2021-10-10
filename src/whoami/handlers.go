package whoami

import (
	"fmt"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/furizu/src/utils"
)

type WhoamiChallenge struct {
	Email string    `dynamo:"email"`
	Token string    `dynamo:"token"`
	Exp   time.Time `dynamo:"exp"`
}

type StartWhoamiReq struct {
	Email string `json:"email"`
}

type RedeemWhoamiReq struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

const CHALLENGES_TABLE = "WhoamiChallenges"

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
	table := utils.FurizuDB.Table(CHALLENGES_TABLE)
	// generate random token
	token := utils.RandomString(10)
	fmt.Printf("token: %s", token)
	// set exp time
	exp := time.Now().Add(time.Hour * 1)
	// TODO: some way to prevent one email from creating multiple reqs
	// generate payload
	payload := WhoamiChallenge{
		Email: req.Email,
		Token: token,
		Exp:   exp,
	}
	// save to db
	err := table.Put(payload).Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
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
	var result WhoamiChallenge
	table := utils.FurizuDB.Table(CHALLENGES_TABLE)
	err := table.Get("token", req.Token).One(&result)
	if err != nil && strings.Contains(err.Error(), "no item found") {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Token does not exist"})
		return
	}

	if err != nil {
		fmt.Printf("err %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// validate
	expired := result.Exp.Before(time.Now())
	invalidEmail := result.Email != req.Email
	// return 400 if invalid
	if expired || invalidEmail {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Token invalid or expired"})
		return
	}
	// TODO: issue JWT, delete token from dynamo
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
