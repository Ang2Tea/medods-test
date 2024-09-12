package common

import "os"

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

const (
	POSTGRES_USERNAME      = "POSTGRES_USERNAME"
	POSTGRES_PASSWORD      = "POSTGRES_PASSWORD"
	POSTGRES_HOST          = "POSTGRES_HOST"
	POSTGRES_PORT          = "POSTGRES_PORT"
	POSTGRES_DATABASE_NAME = "POSTGRES_DATABASE_NAME"
)

func LookupEnv(out *string, key string, defaultVal ...string) {
	val, exist := os.LookupEnv(key)

	if !exist {
		if len(defaultVal) > 0 {
			*out = defaultVal[0]
			return
		} else {
			panic(key + " not found in .env variable")
		}
	}

	*out = val
}
