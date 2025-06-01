package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type requestParams struct {
	Body string `json:"body"`
}

type returnVals struct {
	Error string `json:"error"`
	Valid bool   `json:"valid"`
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	resVals := returnVals{}

	decoder := json.NewDecoder(r.Body)
	reqParams := requestParams{}
	err := decoder.Decode(&reqParams)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		resVals.Error = "Something went wrong"
	}

	if len(reqParams.Body) <= 140 {
		resVals.Valid = true
		w.WriteHeader(200)
	} else {
		resVals.Error = "Chirp is too long"
		w.WriteHeader(400)
	}

	dat, err := json.Marshal(resVals)
	if err != nil {
		log.Printf("Error encoding parameters: %s", err)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dat)
}
