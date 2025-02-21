package pets_crud_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyuff/es-commands"
	"github.com/kyuff/example-petstore/internal/assert"
	"github.com/kyuff/example-petstore/internal/domain/pets"
	"github.com/kyuff/example-petstore/internal/domain/values"
	"github.com/kyuff/example-petstore/internal/features/pets_crud"
)

func TestNewAddPetHandler(t *testing.T) {
	t.Run("fail with invalid content", func(t *testing.T) {
		// arrange
		var (
			dispatcher = &DispatcherMock{}
			sut        = pets_crud.NewAddPetHandler(dispatcher)
			r          = httptest.NewRequest(http.MethodPost, "/pets", bytes.NewBufferString(`{invalid'`))
			w          = httptest.NewRecorder()
		)

		// act
		sut(w, r)

		// assert
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("fail with dispatcher", func(t *testing.T) {
		// arrange
		var (
			dispatcher = &DispatcherMock{}
			sut        = pets_crud.NewAddPetHandler(dispatcher)
			r          = httptest.NewRequest(http.MethodPost, "/pets", bytes.NewBufferString(`{}`))
			w          = httptest.NewRecorder()
		)

		dispatcher.DispatchFunc = func(ctx context.Context, entityID string, cmd commands.Command) error {
			return errors.New("dispatch error")
		}

		// act
		sut(w, r)

		// assert
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
		assert.Equal(t, 1, len(dispatcher.DispatchCalls()))
	})

	t.Run("dispatch add pet command", func(t *testing.T) {
		// arrange
		var (
			dispatcher = &DispatcherMock{}
			body       = `
{
	"id" : 5,
	"name" : "Jack",
	"category" : { "id": 7, "name" : "Category A" },
	"photoUrls" : ["url 1", "url 2"],
	"status" : "test status",
	"tags" : [{ "id": 3, "name": "Tag 3" }, { "id": 4, "name": "Tag 4" }]
}`
			sut = pets_crud.NewAddPetHandler(dispatcher)
			r   = httptest.NewRequest(http.MethodPost, "/pets", bytes.NewBufferString(body))
			w   = httptest.NewRecorder()
		)

		dispatcher.DispatchFunc = func(ctx context.Context, entityID string, cmd commands.Command) error {
			return nil
		}

		// act
		sut(w, r)

		// assert
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
		if assert.Equal(t, 1, len(dispatcher.DispatchCalls())) {
			assert.Type(t, dispatcher.DispatchCalls()[0].Cmd, func(got pets.AddPetCommand) {
				assert.Equal(t, 5, got.ID)
				assert.Equal(t, "Jack", got.PetName)
				assert.Equal(t, values.Category{ID: 7, Name: "Category A"}, got.Category)
				assert.EqualSlice(t, []values.PhotoURL{"url 1", "url 2"}, got.PhotoURLs)
				assert.Equal(t, "TEST STATUS", got.Status)
				assert.EqualSlice(t, []values.Tag{{ID: 3, Name: "Tag 3"}, {ID: 4, Name: "Tag 4"}}, got.Tags)
			})
		}
	})
}
