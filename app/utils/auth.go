package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/furizu/app/users"
)

func Authenticate(c *gin.Context) (user *users.User, err error) {
	if user, ok := c.Get(users.USERS_CTX_KEY); ok {
		return user.(*users.User), nil
	}
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		return nil, errors.New("no token")
	}

	decoded, err := FurizuJWT.ValidateFromToken(token)
	if err != nil {
		return nil, err
	}

	userId, err := decoded.Get("email")
	if err != nil {
		return nil, err
	}
	user, err = users.FindUserByEmail(userId.(string))
	if err != nil {
		return nil, err
	}
	c.Set(users.USERS_CTX_KEY, user)
	return user, nil
}
