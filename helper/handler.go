package helper

import (
	"encoding/json"
	"net/http"
	res "victoria-falls/pkg/response"
)

func SendResponse(w http.ResponseWriter, response *res.Response) {
	w.Header().Add("Content-Type", "application/json")
	code, data := res.Result(response)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetRequest(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
