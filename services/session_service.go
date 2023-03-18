package services

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GenerateSessionService(req requests.GenerateSessionRequest) responses.GenerateSessionResponse {
	resp := responses.NewGenerateSessionResponse()
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
		sessionID := generateSessionID()
		err := repository.DeleteSessionByUserID(user.UUID)
		if err != nil {
			resp.Error = constants.InternalServerError
			return resp
		}
		session, err := repository.CreateSession(user.UUID, sessionID)
		if err != nil {
			fmt.Println("Here &&&&: ", err.Error())
			resp.Error = constants.InternalServerError
			return resp
		}
		resp.AddAllHeaders()
		resp.AddSessionCookie(user.Username, session.SessionID)
		return resp
	default:
		resp.Error = constants.InternalServerError
		return resp
	}
}
