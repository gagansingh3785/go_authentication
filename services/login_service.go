package services

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func LoginService(req requests.LoginRequest) responses.LoginResponse {
	resp := responses.LoginResponse{}

	username := req.Username
	user, err := repository.GetUserByEmailOrUsername(username)
	fmt.Printf("\n %+v \n", user)
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.InvalidCredentials
		resp.Message = constants.InvalidCredentials
	case nil:
		resp.AddAllHeaders()
		resp.Salt = user.Salt
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}
