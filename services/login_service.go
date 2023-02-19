package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func LoginService(req requests.LoginRequest) responses.LoginResponse {
	resp := responses.LoginResponse{}

	username := req.Username
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		user, err = repository.GetUserByEmail(req.Username)
	}
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.InvalidCredentials
	case nil:
		resp.AddAllHeaders()
		resp.Salt = user.Salt
		resp.PasswordHash = user.PasswordHash
	default:
		resp.Error = constants.InternalServerError
	}
	return resp
}
