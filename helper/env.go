package helper

import "os"

// GetEnv - Get Env from .env file
func GetEnv(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
