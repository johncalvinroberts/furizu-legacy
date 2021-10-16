package whoami

import (
	"time"

	"github.com/johncalvinroberts/furizu/src/utils"
)

const CHALLENGES_TABLE = "WhoamiChallenges"

type WhoamiChallenge struct {
	Email string    `dynamo:"email"`
	Token string    `dynamo:"token"` // primary key
	Exp   time.Time `dynamo:"exp"`
}

func upsertWhoamiChallenge(email string) (token string, err error) {
	table := utils.FurizuDB.Table(CHALLENGES_TABLE)
	token = utils.RandomString(10)
	// generate random token
	// set exp time
	exp := time.Now().Add(time.Hour * 1)
	// TODO: some way to prevent one email from creating multiple reqs
	// generate payload
	payload := WhoamiChallenge{
		Email: email,
		Token: token,
		Exp:   exp,
	}
	err = table.Put(payload).Run()
	return token, err
}

func findWhoamiChallenge(token string) (result *WhoamiChallenge, err error) {
	table := utils.FurizuDB.Table(CHALLENGES_TABLE)
	err = table.Get("token", token).One(&result)
	return result, err
}
