package users

import (
	"log"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
	"github.com/johncalvinroberts/furizu/src/utils"
)

const USERS_TABLE = "Users"
const USERS_UUID_NS = "3033d171-09f3-4648-8c28-843e73a5b7e7"

var table dynamo.Table
var userUuidNamespace uuid.UUID

type User struct {
	Email        string    `dynamo:"email"`
	Id           string    `dynamo:"id"`
	CreatedAt    time.Time `dynamo:"createdAt"`
	lastUpsertAt time.Time `dynamo:"lastUpsertAt"`
}

func init() {
	table = utils.FurizuDB.Table(USERS_TABLE)
	res, err := uuid.FromString(USERS_UUID_NS)

	if err != nil {
		log.Fatalf("failed to parse UUID: %v", err)
	}
	userUuidNamespace = res
}

func UpsertUser(email string) (user *User, err error) {
	user, err = FindUserByEmail(email)
	isNotFound := strings.Contains(err.Error(), "no item found")

	if err != nil && !isNotFound {
		log.Fatalf("Failed to query for user %v", err)
		return nil, err
	}

	if err != nil && isNotFound {
		user = &User{
			Email:        email,
			Id:           uuid.NewV5(userUuidNamespace, email).String(),
			CreatedAt:    time.Now(),
			lastUpsertAt: time.Now(),
		}
	}
	if !isNotFound {
		user.lastUpsertAt = time.Now()
	}
	err = table.Put(user).Run()
	if err != nil && !isNotFound {
		log.Fatalf("Failed to put %v", err)
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(email string) (user *User, err error) {
	err = table.Get("email", email).One(&user)
	return user, err
}
