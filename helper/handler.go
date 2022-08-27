package helper

import (
	res "bread-api/pkg/response"
	"encoding/json"
	"net/http"
)

// SendResponse - Return Response
func SendResponse(w http.ResponseWriter, response *res.Response) {
	w.Header().Add("Content-Type", "application/json")
	code, data := res.Result(response)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Get Request - Get Request from Body
func GetRequest(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
