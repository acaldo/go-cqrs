package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/acaldo/cqrs/events"
	"github.com/acaldo/cqrs/models"
	"github.com/acaldo/cqrs/repository"
	"github.com/segmentio/ksuid"
)

type createFeedRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createdFeedHandler(w http.ResponseWriter, r *http.Request) {
	var req createFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	feed := &models.Feed{
		ID:          id.String(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   createdAt,
	}

	if err := repository.InsertFeed(r.Context(), feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Publish new feed event
	if err := events.PublishCreatedFeed(context.Background(), feed); err != nil {
		log.Printf("Error publishing event: %v", err) // Aunque el publish no se invoca, se guarda bd
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}
