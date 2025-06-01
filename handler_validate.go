package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	bodyTextSlice := strings.Split(params.Body, " ")
	prohibitedWords := []string{"kerfuffle", "sharbert", "fornax"}

	cleanedBodyTextSlice := []string{}
	for _, word := range bodyTextSlice {
		wordLower := strings.ToLower(word)
		if slices.Contains(prohibitedWords, wordLower) {
			cleanedBodyTextSlice = append(cleanedBodyTextSlice, "****")
		} else {
			cleanedBodyTextSlice = append(cleanedBodyTextSlice, word)
		}
	}
	cleanedBodyText := strings.Join(cleanedBodyTextSlice, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBodyText,
	})
}
