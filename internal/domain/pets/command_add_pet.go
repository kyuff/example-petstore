package pets

import (
	"context"

	"github.com/kyuff/es"
	"github.com/kyuff/example-petstore/internal/domain/values"
)

type AddPetCommand struct {
	ID        values.PetID      `json:"id"`
	PetName   values.PetName    `json:"name"`
	PhotoURLs []values.PhotoURL `json:"photo_urls"`
	Category  values.Category   `json:"category"`
	Status    values.PetStatus  `json:"status"`
	Tags      []values.Tag      `json:"tags"`
}

func (cmd AddPetCommand) CommandName() string {
	return "AddPetCommand"

}

func (cmd AddPetCommand) Validate() error {
	return nil
}

func NewAddPetCommandExecutor() func(ctx context.Context, cmd AddPetCommand, state *State) ([]es.Content, error) {
	return func(ctx context.Context, cmd AddPetCommand, state *State) ([]es.Content, error) {
		// TODO handle idempotence and general business logic
		return []es.Content{
			Added{},
		}, nil
	}
}
