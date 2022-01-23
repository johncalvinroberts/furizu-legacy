package users

import (
	"log"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
)

const USERS_UUID_NS = "3033d171-09f3-4648-8c28-843e73a5b7e7"

var table dynamo.Table
var userUuidNamespace uuid.UUID

type User struct {
	Email        string    `dynamo:"email"` //primary key
	Id           string    `dynamo:"id"`
	CreatedAt    time.Time `dynamo:"createdAt"`
	LastUpsertAt time.Time `dynamo:"lastUpsertAt"`
}

func InitRepository(db *dynamo.DB, tableName string) {
	table = db.Table(tableName)
	res, err := uuid.FromString(USERS_UUID_NS)

	if err != nil {
		log.Fatalf("failed to parse UUID: %v", err)
	}
	userUuidNamespace = res
}

func UpsertUser(email string) (user *User, err error) {
	user, err = FindUserByEmail(email)
	isNotFound := err != nil && strings.Contains(err.Error(), "no item found")

	if err != nil && !isNotFound {
		log.Fatalf("Failed to query for user %v", err)
		return nil, err
	}

	if isNotFound {
		user = &User{
			Email:        email,
			Id:           uuid.NewV5(userUuidNamespace, email).String(),
			CreatedAt:    time.Now(),
			LastUpsertAt: time.Now(),
		}
		err = table.Put(user).Run()
		return user, err
	}
	// user is a returning user
	err = table.Update("email", email).Set("lastUpsertAt", time.Now()).Run()
	return user, err
}

func FindUserByEmail(email string) (user *User, err error) {
	err = table.Get("email", email).One(&user)
	return user, err
}

func FindUserById(id string) (user *User, err error) {
	err = table.Get("id", id).One(&user)
	return user, err
}
