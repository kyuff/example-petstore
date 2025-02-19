package app

import (
	"errors"

	"github.com/kyuff/anchor"
	"github.com/kyuff/es-commands"
	"github.com/kyuff/example-petstore/internal/domain/pets"
)

var dispatcher = anchor.Singleton(func() (*commands.Dispatcher, error) {
	return commands.NewDispatcher(
		eventStore(),
		commands.SLogMiddleware(logger()),
	), nil
})

func registerCommands() error {
	return errors.Join(
		commands.RegisterFunc(dispatcher(), pets.EntityType, pets.NewAddPetCommandExecutor()),
	)
}
