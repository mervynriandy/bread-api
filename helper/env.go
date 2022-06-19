package helper

import "os"

// envString function receives env key string
// If env file doesn't contain the requested key
// returns fallback string
func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
