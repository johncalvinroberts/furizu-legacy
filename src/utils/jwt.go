package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	jwtpkg "github.com/robbert229/jwt"
)

type JWTEncDec struct {
	alg jwtpkg.Algorithm
}

var FurizuJWT *JWTEncDec

var ttlMs int

func InitJWT() {
	FurizuJWT = &JWTEncDec{
		alg: jwtpkg.HmacSha256(os.Getenv("JWT_SECRET")),
	}
	var err error
	ttlMs, err = strconv.Atoi(os.Getenv("JWT_TTL_MINS"))
	ttlMs = ttlMs * 60000
	if err != nil {
		log.Fatalf("Failed to init JWT")
	}
}

func (ed *JWTEncDec) FromToken(token string, kvs map[string]string) (map[string]string, error) {
	claims, err := ed.alg.Decode(token)
	if err != nil {
		return nil, err
	}

	for key := range kvs {
		iVal, err := claims.Get(key)
		if err != nil {
			return nil, err
		}
		strVal, ok := iVal.(string)
		if !ok {
			return nil, errors.New("incorrect JWT claim")
		}

		kvs[key] = strVal
	}
	return kvs, nil
}

func (ed *JWTEncDec) ToToken(kvs map[string]string) (string, error) {
	claims := jwtpkg.NewClaim()
	claims.Set("expire", time.Now().Unix()+int64(ttlMs))
	for key, val := range kvs {
		claims.Set(key, val)
	}

	token, err := ed.alg.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
