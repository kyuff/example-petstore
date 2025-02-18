package app

import (
	"github.com/kyuff/anchor"
	"github.com/kyuff/es"
	"github.com/kyuff/es/storage/inmemory"
)

var eventStore = anchor.Singleton(func() (*es.Store, error) {
	return es.NewStore(
		// TODO Upgrade to a persistent storage
		inmemory.New(),
		es.WithSlog(logger()),
		es.WithEvents(
			"TODO",
			// TODO Register events
			nil,
		),
	), nil
})
