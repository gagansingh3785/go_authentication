package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func RegisterService(registerRequest requests.RegisterRequest) responses.RegisterResponse {
	resp := responses.RegisterResponse{}
	username := registerRequest.Username
	salt := registerRequest.Salt
	phone := registerRequest.Phone
	passwordHash := registerRequest.PasswordHash
	email := registerRequest.Email
	_, err := repository.CreateNewUser(username, salt, phone, passwordHash, email, 1)
	switch err {
	case nil:
		resp.Error = ""
		resp.Message = "Registration Successful :)"
		resp.AddAllHeaders()
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}
