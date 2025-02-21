package petstore

import (
	"net/http"

	"github.com/kyuff/example-petstore/generated/api"
)

type API struct {
	AddPetFunc                   func(w http.ResponseWriter, r *http.Request)
	UpdatePetFunc                func(w http.ResponseWriter, r *http.Request)
	FindPetsByStatusFunc         func(w http.ResponseWriter, r *http.Request, params api.FindPetsByStatusParams)
	FindPetsByTagsFunc           func(w http.ResponseWriter, r *http.Request, params api.FindPetsByTagsParams)
	DeletePetFunc                func(w http.ResponseWriter, r *http.Request, petId int64, params api.DeletePetParams)
	GetPetByIdFunc               func(w http.ResponseWriter, r *http.Request, petId int64)
	UpdatePetWithFormFunc        func(w http.ResponseWriter, r *http.Request, petId int64, params api.UpdatePetWithFormParams)
	UploadFileFunc               func(w http.ResponseWriter, r *http.Request, petId int64, params api.UploadFileParams)
	GetInventoryFunc             func(w http.ResponseWriter, r *http.Request)
	PlaceOrderFunc               func(w http.ResponseWriter, r *http.Request)
	DeleteOrderFunc              func(w http.ResponseWriter, r *http.Request, orderId int64)
	GetOrderByIdFunc             func(w http.ResponseWriter, r *http.Request, orderId int64)
	CreateUserFunc               func(w http.ResponseWriter, r *http.Request)
	CreateUsersWithListInputFunc func(w http.ResponseWriter, r *http.Request)
	LoginUserFunc                func(w http.ResponseWriter, r *http.Request, params api.LoginUserParams)
	LogoutUserFunc               func(w http.ResponseWriter, r *http.Request)
	DeleteUserFunc               func(w http.ResponseWriter, r *http.Request, username string)
	GetUserByNameFunc            func(w http.ResponseWriter, r *http.Request, username string)
	UpdateUserFunc               func(w http.ResponseWriter, r *http.Request, username string)
}

var _ api.ServerInterface = (*API)(nil)

func (api *API) AddPet(w http.ResponseWriter, r *http.Request) {
	api.AddPetFunc(w, r)
}

func (api *API) UpdatePet(w http.ResponseWriter, r *http.Request) {
	api.UpdatePetFunc(w, r)
}

func (api *API) FindPetsByStatus(w http.ResponseWriter, r *http.Request, params api.FindPetsByStatusParams) {
	api.FindPetsByStatusFunc(w, r, params)
}

func (api *API) FindPetsByTags(w http.ResponseWriter, r *http.Request, params api.FindPetsByTagsParams) {
	api.FindPetsByTagsFunc(w, r, params)
}

func (api *API) DeletePet(w http.ResponseWriter, r *http.Request, petId int64, params api.DeletePetParams) {
	api.DeletePetFunc(w, r, petId, params)
}

func (api *API) GetPetById(w http.ResponseWriter, r *http.Request, petId int64) {
	api.GetPetByIdFunc(w, r, petId)
}

func (api *API) UpdatePetWithForm(w http.ResponseWriter, r *http.Request, petId int64, params api.UpdatePetWithFormParams) {
	api.UpdatePetWithFormFunc(w, r, petId, params)
}

func (api *API) UploadFile(w http.ResponseWriter, r *http.Request, petId int64, params api.UploadFileParams) {
	api.UploadFileFunc(w, r, petId, params)
}

func (api *API) GetInventory(w http.ResponseWriter, r *http.Request) {
	api.GetInventoryFunc(w, r)
}

func (api *API) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	api.PlaceOrderFunc(w, r)
}

func (api *API) DeleteOrder(w http.ResponseWriter, r *http.Request, orderId int64) {
	api.DeleteOrderFunc(w, r, orderId)
}

func (api *API) GetOrderById(w http.ResponseWriter, r *http.Request, orderId int64) {
	api.GetOrderByIdFunc(w, r, orderId)
}

func (api *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	api.CreateUserFunc(w, r)
}

func (api *API) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	api.CreateUsersWithListInputFunc(w, r)
}

func (api *API) LoginUser(w http.ResponseWriter, r *http.Request, params api.LoginUserParams) {
	api.LoginUserFunc(w, r, params)
}

func (api *API) LogoutUser(w http.ResponseWriter, r *http.Request) {
	api.LogoutUserFunc(w, r)
}

func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request, username string) {
	api.DeleteUserFunc(w, r, username)
}

func (api *API) GetUserByName(w http.ResponseWriter, r *http.Request, username string) {
	api.GetUserByNameFunc(w, r, username)
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request, username string) {
	api.UpdateUserFunc(w, r, username)
}
