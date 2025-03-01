package pkg

import "github.com/google/uuid"

func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
} 