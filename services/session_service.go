package services

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GenerateSessionService(req requests.GenerateSessionRequest) responses.GenerateSessionResponse {
	resp := responses.GenerateSessionResponse{}

	username := req.Username
	passwordHash := req.PasswordHash
	user, err := repository.GetUserByEmailOrUsername(username)
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.InvalidCredentials
		return resp
	case nil:
		if active, err := isLoggedIn(user.UUID); err != nil {
			resp.Error = constants.InternalServerError
			resp.Message = constants.InternalServerError
			return resp
		} else if active {
			resp.Error = constants.AlreadyLoggedIn
			resp.Message = constants.AlreadyLoggedIn
			return resp
		}
		if user.PasswordHash != passwordHash {
			resp.Error = constants.InvalidCredentials
			return resp
		}
		sessionID := generateSessionID()
		session, err := repository.CreateSession(user.UUID, sessionID)
		if err != nil {
			fmt.Println("Here &&&&: ", err.Error())
			resp.Error = constants.InternalServerError
			return resp
		}
		resp.AddAllHeaders(user.Username, session.SessionID)
		if resp.Headers[constants.SESSION_COOKIE] == "" {
			resp.Error = constants.InternalServerError
		}
		return resp
	default:
		resp.Error = constants.InternalServerError
		return resp
	}
}
