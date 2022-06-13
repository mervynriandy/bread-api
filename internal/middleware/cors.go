package middleware

import (
	"net/http"
)

func SetCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methdos", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
}
