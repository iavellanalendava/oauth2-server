package model

import (
	"time"

	"github.com/google/uuid"
)

type Key struct {
	Id        uuid.UUID
	PublicKey []byte
	CreatedAt time.Time

	// Additional fields can be added as needed
}
