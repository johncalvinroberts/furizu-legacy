package utils

import (
	"fmt"
	"os"
)

var localDefaults = map[string]string{
	"JWT_SECRET":       "secretshhhh",
	"JWT_TTL_MINS":     "60",
	"COOKIE_SECURE":    "false",
	"COOKIE_HTTP_ONLY": "false",
	"GIN_MODE":         "debug",
	// tables
	"USERS_TABLE":             "Users",
	"WHOAMI_CHALLENGES_TABLE": "WhoamiChallenges",
}

func InitEnv() {
	fmt.Println("LAOADING ENV")
	for k, v := range localDefaults {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
