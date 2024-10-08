package handlers

import (
	"context"
	"dz_first_pantela/iternal/usersService"
	"dz_first_pantela/iternal/web/users"
)

type UserHandler struct {
	Service *usersService.UserService
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	result, err := h.Service.DeleteUserById(int(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId200JSONResponse(result), nil
}

func (h *UserHandler) GetUsers(_ context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, msg := range allUsers {
		user := users.User{
			Id:       &msg.ID,
			Email:    &msg.Email,
			Password: &msg.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := usersService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PutUsersId(_ context.Context, request users.PutUsersIdRequestObject) (users.PutUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body
	userToUpdate := usersService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserById(int(id), userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PutUsersId201JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func NewUserHandler(service *usersService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
