package services

import (
	"fmt"
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
	if present, err := isDuplicateIdentifiePresent(username); err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	} else if present {
		resp.Error = constants.UsernameTaken
		resp.Message = constants.UsernameTaken
		return resp
	}
	if present, err := isDuplicateIdentifiePresent(email); err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	} else if present {
		resp.Error = constants.EmailAlreadyTaken
		resp.Message = constants.EmailAlreadyTaken
		return resp
	}
	user, err := repository.CreateNewUser(username, email, salt, phone, passwordHash, 1)
	switch err {
	case nil:
		fmt.Printf("\n %+v \n", user)
		resp.Error = ""
		resp.Message = "Registration Successful :)"
		resp.AddAllHeaders()
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

func isDuplicateIdentifiePresent(identifier string) (bool, error) {
	_, err := repository.GetUserByEmailOrUsername(identifier)
	if err == constants.ErrSQLNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
