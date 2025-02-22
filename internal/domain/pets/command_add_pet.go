package pets

import (
	"context"

	"github.com/kyuff/es"
	commands "github.com/kyuff/es-commands"
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

func NewAddPetCommandExecutor() commands.ExecutorFunc[AddPetCommand, *State] {
	return func(ctx context.Context, cmd AddPetCommand, state *State) ([]es.Content, error) {
		if state.ID == cmd.ID {
			// idempotency
			return nil, nil
		}
		return []es.Content{
			AddedV1{
				ID:      cmd.ID,
				PetName: cmd.PetName,
			},
		}, nil
	}
}
