package app

import (
	"errors"

	"github.com/kyuff/anchor"
	commands "github.com/kyuff/es-commands"
	"github.com/kyuff/example-petstore/internal/domain/pets"
	"github.com/kyuff/validate"
)

var dispatcher = anchor.Singleton(func() (*commands.Dispatcher, error) {
	return commands.NewDispatcher(
		eventStore(),
		commands.SLogMiddleware(logger()),
		commands.Validate(validate.Middleware),
	), nil
})

func registerCommands() error {
	return errors.Join(
		commands.RegisterFunc(dispatcher(), pets.EntityType, pets.NewAddPetCommandExecutor()),
	)
}
