package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Status  int
	Message string
	Body    map[string]string
}

func ReturnAPIResponse(w http.ResponseWriter, r *http.Request, status int, message string, body map[string]string) {
	response := ResponseMessage{
		Status:  status,
		Message: message,
		Body:    body,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
