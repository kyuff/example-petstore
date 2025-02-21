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
	case Added:
		s.applyAdded(e)
	}
	return nil
}

func (s *State) applyAdded(e Added) {
	s.ID = e.ID
}
