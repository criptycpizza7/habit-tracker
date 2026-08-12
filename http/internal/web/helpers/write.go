package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func WriteResponse(w http.ResponseWriter, status int, data any) {
	resp := response{
		Status: "success",
		Data:   data,
	}
	var raw_resp []byte
	if data != nil {
		var err error
		raw_resp, err = json.Marshal(resp)
		if err != nil {
			log.Println(err)
			WriteEmptyError(w, http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(status)
	_, err := w.Write(raw_resp)
	if err != nil {
		log.Println(err)
		WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	resp := response{
		Status: "error",
		Data:   message,
	}
	var raw_resp []byte
	if message != "" {
		var err error
		raw_resp, err = json.Marshal(resp)
		if err != nil {
			WriteEmptyError(w, http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(status)
	_, err := w.Write(raw_resp)
	if err != nil {
		WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
}

func WriteEmptyError(w http.ResponseWriter, status int) {
	resp := response{
		Status: "error",
		Data:   nil,
	}
	raw_resp, _ := json.Marshal(resp)
	w.WriteHeader(status)
	w.Write(raw_resp)
}
