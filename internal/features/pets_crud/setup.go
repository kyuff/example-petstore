package pets_crud

import (
	"context"

	"github.com/kyuff/es-commands"
	"github.com/kyuff/example-petstore/internal/petstore"
)

type Dispatcher interface {
	Dispatch(ctx context.Context, entityID string, cmd commands.Command) error
}

func New(api *petstore.API, dispatcher Dispatcher) error {
	api.AddPetFunc = NewAddPetHandler(dispatcher)
	return nil
}
