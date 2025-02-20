package pets_crud

import (
	"context"
	"fmt"
	"strings"

	"github.com/kyuff/example-petstore/generated/api"
	"github.com/kyuff/example-petstore/internal/domain/pets"
	"github.com/kyuff/example-petstore/internal/domain/values"
)

func NewAddPetHandler(dispatcher Dispatcher) func(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error) {
	mapStatus := func(s *api.PetStatus) values.PetStatus {
		if s == nil {
			return ""
		}

		return values.PetStatus(strings.ToUpper(string(*s)))
	}

	mapPhotoURLs := func(urls []string) []values.PhotoURL {
		var result []values.PhotoURL
		for _, url := range urls {
			result = append(result, values.PhotoURL(url))
		}

		return result
	}

	mapCategory := func(category *api.Category) values.Category {
		if category == nil {
			return values.Category{}
		}

		return values.Category{
			ID:   nonNil(category.Id),
			Name: nonNil(category.Name),
		}
	}

	mapTags := func(tags *[]api.Tag) []values.Tag {
		if tags == nil {
			return nil
		}

		var result []values.Tag
		for _, tag := range *tags {
			result = append(result, values.Tag{
				ID:   nonNil(tag.Id),
				Name: nonNil(tag.Name),
			})
		}

		return result
	}

	return func(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error) {
		var pet api.Pet
		var ok api.AddPetResponseObject
		if request.JSONBody != nil {
			pet = *request.JSONBody
			ok = api.AddPet200JSONResponse(pet)
		} else if request.FormdataBody != nil {
			pet = *request.FormdataBody
			ok = api.AddPet200JSONResponse(pet)
		} else {
			return nil, fmt.Errorf("no request provided")
		}

		var petID = values.PetIDFromPtr(pet.Id)
		err := dispatcher.Dispatch(ctx, petID.String(), pets.AddPetCommand{
			ID:        petID,
			PetName:   values.PetName(pet.Name),
			PhotoURLs: mapPhotoURLs(pet.PhotoUrls),
			Category:  mapCategory(pet.Category),
			Status:    mapStatus(pet.Status),
			Tags:      mapTags(pet.Tags),
		})
		if err != nil {
			return nil, err
		}

		return ok, nil
	}
}

func nonNil[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}

	var t T
	return t
}
