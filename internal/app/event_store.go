package app

import (
	"github.com/kyuff/anchor"
	"github.com/kyuff/es"
	"github.com/kyuff/es/storage/inmemory"
	"github.com/kyuff/example-petstore/internal/domain/pets"
)

var eventStore = anchor.Singleton(func() (*es.Store, error) {
	return es.NewStore(
		// TODO Upgrade to a persistent storage
		inmemory.New(),
		es.WithSlog(logger()),
		es.WithEvents(pets.EntityType, pets.Events),
	), nil
})
