package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GenerateSessionService(req requests.GenerateSessionRequest) responses.GenerateSessionResponse {
	resp := responses.GenerateSessionResponse{}
	resp.Headers = make(map[string]string)

	username := req.Username
	passwordHash := req.PasswordHash
	user, err := repository.GetUserByEmailOrUsername(username)
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.InvalidCredentials
		return resp
	case nil:
		if user.PasswordHash != passwordHash {
			resp.Error = constants.InvalidCredentials
			return resp
		}
		sessionID := getSessionID()
		if err := repository.CreateSession(user.UUID, sessionID); err != nil {
			resp.Error = constants.InternalServerError
			return resp
		}
		resp.Message = "Login Successful"
		resp.AddAllHeaders(sessionID)
		return resp
	default:
		resp.Error = constants.InternalServerError
		return resp
	}
}

func getSessionID() string {
	return ""
}
