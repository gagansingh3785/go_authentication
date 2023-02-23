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
	case nil:
		if loggedIn, err := isLoggedIn(user.UUID); err != nil {
			resp.Error = constants.InternalServerError
			resp.Message = constants.InternalServerError
			return resp
		} else if loggedIn {
			resp.Error = constants.AlreadyLoggedIn
			resp.Message = constants.AlreadyLoggedIn
			return resp
		}
		resp.AddAllHeaders()
		resp.Salt = user.Salt
		resp.PasswordHash = user.PasswordHash
	default:
		resp.Error = constants.InternalServerError
	}
	return resp
}

func isLoggedIn(userID string) (bool, error) {
	_, err := repository.GetSessionFromUserID(userID)
	switch err {
	case constants.ErrSQLNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
