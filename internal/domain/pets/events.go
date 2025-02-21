package pets

import (
	"github.com/kyuff/es"
	"github.com/kyuff/example-petstore/internal/domain/values"
)

var EntityType = "pets"
var Events = []es.Content{
	Added{},
}

type Added struct {
	ID values.PetID `json:"id"`
}

func (e Added) Name() string {
	return "AddedV1"
}
