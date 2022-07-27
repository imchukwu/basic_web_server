package api

import (
	"github.com/google/uuid"
)

type Item struct {
	id   uuid.UUID `json:"id"`
	name string    `json:"name"`
}
