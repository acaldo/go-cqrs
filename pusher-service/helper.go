package main

import "time"

type CreatedFeedMessage struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewCreatedFeedMessage(id string, title string, description string, createdAt time.Time) *CreatedFeedMessage {
	return &CreatedFeedMessage{
		Type:        "created_feed",
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
