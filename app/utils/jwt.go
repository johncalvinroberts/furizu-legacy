package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	jwtpkg "github.com/robbert229/jwt"
)

const EXPIRE_CLAIM = "expire"

type jwtEncDec struct {
	alg jwtpkg.Algorithm
}

var FurizuJWT *jwtEncDec

var ttlMs int

func InitJWT() {
	FurizuJWT = &jwtEncDec{
		alg: jwtpkg.HmacSha256(os.Getenv("JWT_SECRET")),
	}
	var err error
	ttlMs, err = strconv.Atoi(os.Getenv("JWT_TTL_MINS"))
	ttlMs = ttlMs * 60000
	if err != nil {
		log.Fatalf("Failed to init JWT")
	}
}

func (ed *jwtEncDec) FromToken(token string, kvs map[string]string) (map[string]string, error) {
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

func (ed *jwtEncDec) ToToken(kvs map[string]string) (string, error) {
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
