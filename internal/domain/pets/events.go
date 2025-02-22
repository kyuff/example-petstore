package pets

import (
	"github.com/kyuff/es"
	"github.com/kyuff/example-petstore/internal/domain/values"
)

var EntityType = "pets"
var Events = []es.Content{
	AddedV1{},
}

type AddedV1 struct {
	ID      values.PetID   `json:"id"`
	PetName values.PetName `json:"name"`
}

func (e AddedV1) EventName() string {
	return "AddedV1"
}
