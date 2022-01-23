package whoami

import (
	"time"

	"github.com/guregu/dynamo"
	"github.com/johncalvinroberts/furizu/utils"
)

var table dynamo.Table

type WhoamiChallenge struct {
	Email string    `dynamo:"email"`
	Token string    `dynamo:"token"` // primary key
	Exp   time.Time `dynamo:"exp"`
}

func InitRepository(db *dynamo.DB, tableName string) {
	table = utils.FurizuDB.Table(tableName)
}

func upsertWhoamiChallenge(email string) (token string, err error) {
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
	err = table.Get("token", token).One(&result)
	return result, err
}

func destroyWhoamiChallenge(token string) (err error) {
	err = table.Delete("token", token).Run()
	return err
}
