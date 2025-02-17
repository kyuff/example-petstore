package petstore

import (
	"context"

	"github.com/kyuff/example-petstore/generated/api"
)

type API struct {
	AddPetFunc                   func(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error)
	UpdatePetFunc                func(ctx context.Context, request api.UpdatePetRequestObject) (api.UpdatePetResponseObject, error)
	FindPetsByStatusFunc         func(ctx context.Context, request api.FindPetsByStatusRequestObject) (api.FindPetsByStatusResponseObject, error)
	FindPetsByTagsFunc           func(ctx context.Context, request api.FindPetsByTagsRequestObject) (api.FindPetsByTagsResponseObject, error)
	DeletePetFunc                func(ctx context.Context, request api.DeletePetRequestObject) (api.DeletePetResponseObject, error)
	GetPetByIdFunc               func(ctx context.Context, request api.GetPetByIdRequestObject) (api.GetPetByIdResponseObject, error)
	UpdatePetWithFormFunc        func(ctx context.Context, request api.UpdatePetWithFormRequestObject) (api.UpdatePetWithFormResponseObject, error)
	UploadFileFunc               func(ctx context.Context, request api.UploadFileRequestObject) (api.UploadFileResponseObject, error)
	GetInventoryFunc             func(ctx context.Context, request api.GetInventoryRequestObject) (api.GetInventoryResponseObject, error)
	PlaceOrderFunc               func(ctx context.Context, request api.PlaceOrderRequestObject) (api.PlaceOrderResponseObject, error)
	DeleteOrderFunc              func(ctx context.Context, request api.DeleteOrderRequestObject) (api.DeleteOrderResponseObject, error)
	GetOrderByIdFunc             func(ctx context.Context, request api.GetOrderByIdRequestObject) (api.GetOrderByIdResponseObject, error)
	CreateUserFunc               func(ctx context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error)
	CreateUsersWithListInputFunc func(ctx context.Context, request api.CreateUsersWithListInputRequestObject) (api.CreateUsersWithListInputResponseObject, error)
	LoginUserFunc                func(ctx context.Context, request api.LoginUserRequestObject) (api.LoginUserResponseObject, error)
	LogoutUserFunc               func(ctx context.Context, request api.LogoutUserRequestObject) (api.LogoutUserResponseObject, error)
	DeleteUserFunc               func(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error)
	GetUserByNameFunc            func(ctx context.Context, request api.GetUserByNameRequestObject) (api.GetUserByNameResponseObject, error)
	UpdateUserFunc               func(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error)
}

var _ api.StrictServerInterface = (*API)(nil)

func (api *API) AddPet(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error) {
	return api.AddPetFunc(ctx, request)
}

func (api *API) UpdatePet(ctx context.Context, request api.UpdatePetRequestObject) (api.UpdatePetResponseObject, error) {
	return api.UpdatePetFunc(ctx, request)
}

func (api *API) FindPetsByStatus(ctx context.Context, request api.FindPetsByStatusRequestObject) (api.FindPetsByStatusResponseObject, error) {
	return api.FindPetsByStatusFunc(ctx, request)
}

func (api *API) FindPetsByTags(ctx context.Context, request api.FindPetsByTagsRequestObject) (api.FindPetsByTagsResponseObject, error) {
	return api.FindPetsByTagsFunc(ctx, request)
}

func (api *API) DeletePet(ctx context.Context, request api.DeletePetRequestObject) (api.DeletePetResponseObject, error) {
	return api.DeletePetFunc(ctx, request)
}

func (api *API) GetPetById(ctx context.Context, request api.GetPetByIdRequestObject) (api.GetPetByIdResponseObject, error) {
	return api.GetPetByIdFunc(ctx, request)
}

func (api *API) UpdatePetWithForm(ctx context.Context, request api.UpdatePetWithFormRequestObject) (api.UpdatePetWithFormResponseObject, error) {
	return api.UpdatePetWithFormFunc(ctx, request)
}

func (api *API) UploadFile(ctx context.Context, request api.UploadFileRequestObject) (api.UploadFileResponseObject, error) {
	return api.UploadFileFunc(ctx, request)
}

func (api *API) GetInventory(ctx context.Context, request api.GetInventoryRequestObject) (api.GetInventoryResponseObject, error) {
	return api.GetInventoryFunc(ctx, request)
}

func (api *API) PlaceOrder(ctx context.Context, request api.PlaceOrderRequestObject) (api.PlaceOrderResponseObject, error) {
	return api.PlaceOrderFunc(ctx, request)
}

func (api *API) DeleteOrder(ctx context.Context, request api.DeleteOrderRequestObject) (api.DeleteOrderResponseObject, error) {
	return api.DeleteOrderFunc(ctx, request)
}

func (api *API) GetOrderById(ctx context.Context, request api.GetOrderByIdRequestObject) (api.GetOrderByIdResponseObject, error) {
	return api.GetOrderByIdFunc(ctx, request)
}

func (api *API) CreateUser(ctx context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	return api.CreateUserFunc(ctx, request)
}

func (api *API) CreateUsersWithListInput(ctx context.Context, request api.CreateUsersWithListInputRequestObject) (api.CreateUsersWithListInputResponseObject, error) {
	return api.CreateUsersWithListInputFunc(ctx, request)
}

func (api *API) LoginUser(ctx context.Context, request api.LoginUserRequestObject) (api.LoginUserResponseObject, error) {
	return api.LoginUserFunc(ctx, request)
}

func (api *API) LogoutUser(ctx context.Context, request api.LogoutUserRequestObject) (api.LogoutUserResponseObject, error) {
	return api.LogoutUserFunc(ctx, request)
}

func (api *API) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	return api.DeleteUserFunc(ctx, request)
}

func (api *API) GetUserByName(ctx context.Context, request api.GetUserByNameRequestObject) (api.GetUserByNameResponseObject, error) {
	return api.GetUserByNameFunc(ctx, request)
}

func (api *API) UpdateUser(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	return api.UpdateUserFunc(ctx, request)
}
