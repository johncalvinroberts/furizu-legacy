package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/robbert229/jwt"
)

const EXPIRE_CLAIM = "expire"
const USERS_CTX_KEY = "USER"

type jwtEncDec struct {
	alg jwtpkg.Algorithm
}

type UserPartial struct {
	Email string
	Id    string
}

var FurizuJWT *jwtEncDec

var jwtTtlMs int
var refreshTokenTtlMs int

func InitJWT() {
	FurizuJWT = &jwtEncDec{
		alg: jwtpkg.HmacSha256(os.Getenv("JWT_SECRET")),
	}
	var err error
	jwtTtlMs, err = strconv.Atoi(os.Getenv("JWT_ACCESSTOKEN_TTL_MS"))
	if err != nil {
		log.Fatalf("Failed to init JWT")
	}
	refreshTokenTtlMs, err = strconv.Atoi(os.Getenv("JWT_REFRESHTOKEN_TTL_MS"))
	if err != nil {
		log.Fatalf("Failed to init JWT")
	}
}

func (ed *jwtEncDec) CreateToken(kvs map[string]string, ttlMs int) (string, error) {
	claims := jwtpkg.NewClaim()
	claims.Set(EXPIRE_CLAIM, time.Now().Unix()+int64(ttlMs))
	for key, val := range kvs {
		claims.Set(key, val)
	}

	token, err := ed.alg.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ed *jwtEncDec) GenerateAccessToken(kvs map[string]string) (string, error) {
	return ed.CreateToken(kvs, jwtTtlMs)
}

func (ed *jwtEncDec) GenerateRefreshToken(kvs map[string]string) (string, error) {
	return ed.CreateToken(kvs, refreshTokenTtlMs)
}

func (ed *jwtEncDec) ValidateFromToken(token string) (claims *jwtpkg.Claims, err error) {
	err = ed.alg.Validate(token)
	if err != nil {
		return nil, err
	}

	claims, err = ed.alg.DecodeAndValidate(token)
	if err != nil {
		return nil, err
	}
	exp, err := claims.GetTime(EXPIRE_CLAIM)

	if time.Now().Unix() > exp.Unix() {
		return nil, errors.New("token expired")
	}

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func Authenticate(c *gin.Context) (user *UserPartial, err error) {
	if user, ok := c.Get(USERS_CTX_KEY); ok {
		return user.(*UserPartial), nil
	}
	header := c.Request.Header.Get("Authorization")
	token := strings.Split(header, "Bearer")[1]
	token = strings.TrimSpace(token)

	if token == "" {
		return nil, errors.New("no token")
	}

	decoded, err := FurizuJWT.ValidateFromToken(token)
	if err != nil {
		return nil, err
	}

	email, err := decoded.Get("email")
	if err != nil {
		return nil, err
	}
	userId, err := decoded.Get("userId")
	if err != nil {
		return nil, err
	}

	userPartial := &UserPartial{
		Email: email.(string),
		Id:    userId.(string),
	}

	c.Set(USERS_CTX_KEY, userPartial)
	return userPartial, nil
}
