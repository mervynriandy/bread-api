package middleware

import (
	"log"
	"net/http"

	"go.uber.org/zap"
)

type AuthenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *AuthenticationMiddleware) Populate() {
	amw.tokenUsers = make(map[string]string)
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zap.L().Info("middleware", zap.String("routes", r.RequestURI))
		SetCors(w)
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
