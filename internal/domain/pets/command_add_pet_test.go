package pets_test

import (
	"testing"

	"github.com/kyuff/es"
	"github.com/kyuff/example-petstore/internal/assert"
	"github.com/kyuff/example-petstore/internal/domain/pets"
	"github.com/kyuff/example-petstore/internal/domain/values"
	"github.com/kyuff/testdata"
)

func TestNewAddPetCommandExecutor(t *testing.T) {
	var (
		cfg = testdata.NewConfig()
	)

	t.Run("be idempotent", func(t *testing.T) {
		// arrange
		var (
			sut = pets.NewAddPetCommandExecutor()

			_     = testdata.MakeStickyWith[values.PetID](t, cfg)
			added = testdata.MakeWith[pets.AddedV1](t, cfg)
			state = es.HydrateState(t, &pets.State{}, added)
			cmd   = testdata.MakeWith[pets.AddPetCommand](t, cfg)
		)

		// act
		got, err := sut.Execute(t.Context(), cmd, state)

		// assert
		assert.NoError(t, err)
		es.VerifyEvents(t, got)
	})

	t.Run("emit AddedV1", func(t *testing.T) {
		// arrange
		var (
			sut = pets.NewAddPetCommandExecutor()

			state = es.HydrateState(t, &pets.State{})
			cmd   = testdata.MakeWith[pets.AddPetCommand](t, cfg)
		)

		// act
		got, err := sut.Execute(t.Context(), cmd, state)

		// assert
		assert.NoError(t, err)
		es.VerifyEvents(t, got,
			pets.AddedV1{
				ID:      cmd.ID,
				PetName: cmd.PetName,
			},
		)
	})
}
