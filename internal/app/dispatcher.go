package app

import (
	"errors"

	"github.com/kyuff/anchor"
	"github.com/kyuff/example-petstore/internal/domain/pets"
)
import "github.com/kyuff/es-commands"

var dispatcher = anchor.Singleton(func() (*commands.Dispatcher, error) {
	return commands.NewDispatcher(eventStore()), nil
})

func registerCommands() error {
	return errors.Join(
		commands.RegisterFunc(dispatcher(), pets.EntityType, pets.NewAddPetCommandExecutor()),
	)
}
