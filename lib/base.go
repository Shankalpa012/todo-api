package lib

import "github.com/google/uuid"

type ModelBase struct {
	UID uuid.UUID `json:"id"`
}
