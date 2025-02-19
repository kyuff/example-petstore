package pets

import (
	"context"

	"github.com/kyuff/es"
)

type AddPetCommand struct {
}

func (cmd AddPetCommand) Name() string {
	return "AddPetCommand"
}

func NewAddPetCommandExecutor() func(ctx context.Context, cmd AddPetCommand, state *State) ([]es.Content, error) {
	return func(ctx context.Context, cmd AddPetCommand, state *State) ([]es.Content, error) {
		// TODO handle idempotence and general business logic
		return []es.Content{
			Added{},
		}, nil
	}
}
