package pets

import (
	"context"

	"github.com/kyuff/es"
	"github.com/kyuff/example-petstore/internal/domain/values"
)

type State struct {
	ID values.PetID
}

func (s *State) Handle(ctx context.Context, event es.Event) error {
	switch e := event.Content.(type) {
	case AddedV1:
		s.applyAdded(e)
	}
	return nil
}

func (s *State) applyAdded(e AddedV1) {
	s.ID = e.ID
}
