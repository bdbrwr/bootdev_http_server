package main

import (
	"database/sql"
	"net/http"

	"github.com/bdbrwr/bootdev_http_server/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	authorIDString := r.URL.Query().Get("author_id")
	var authorID uuid.NullUUID
	if authorIDString == "" {
		authorID.Valid = false
	} else {
		parsedUUID, err := uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't parse UUID author_id", err)
			return
		}
		authorID.UUID = parsedUUID
		authorID.Valid = true
	}

	var sort sql.NullString
	if r.URL.Query().Get("sort") == "desc" {
		sort.String = "desc"
		sort.Valid = true
	}

	getChirpsParams := database.GetChirpsParams{
		AuthorID: authorID,
		Sort:     sort,
	}

	dbChirps, err := cfg.db.GetChirps(r.Context(), getChirpsParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
