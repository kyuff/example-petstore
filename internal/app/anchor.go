package app

import (
	"errors"

	"github.com/kyuff/anchor"
	"github.com/kyuff/example-petstore/internal/features/pets_crud"
	"github.com/kyuff/example-petstore/internal/features/pets_view"
)

func Run(wire anchor.Wire) int {
	return anchor.New(wire, anchor.WithDefaultSlog()).
		Add(
			anchor.Setup("commands.Dispatcher", registerCommands),
			anchor.Setup("features", func() error {
				return errors.Join(
					pets_crud.New(httpApi(), dispatcher()),
					pets_view.New(httpApi()),
				)
			}),
		).
		Add(
			eventStore(),
			anchor.Make("http.Server", NewHttpServer),
		).
		Run()
}
