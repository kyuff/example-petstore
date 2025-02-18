package pets

import (
	"context"

	"github.com/kyuff/es"
)

type State struct{}

func (s *State) Handle(ctx context.Context, event es.Event) error {
	return nil
}
