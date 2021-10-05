package whoami

import (
	"net/http"

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
	table := utils.FurizuDB.Table("WhoamiChallenges")
	// generate random token
	token := utils.RandomString(10)
	// set exp time
	exp := time.Now().Add(time.Hour * 1)
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

func Redeem(c *gin.Context) {
	// find WhoamiChallenge based on request
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
