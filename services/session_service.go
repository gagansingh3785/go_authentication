package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
	"github.com/google/uuid"
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
		if user.PasswordHash != passwordHash {
			resp.Error = constants.InvalidCredentials
			return resp
		}
		sessionID := generateSessionID()
		session, err := repository.CreateSession(user.UUID, sessionID)
		if err != nil {
			resp.Error = constants.InternalServerError
			return resp
		}
		resp.Message = "Login Successful"
		resp.AddAllHeaders(user.Username, session.SessionID)
		return resp
	default:
		resp.Error = constants.InternalServerError
		return resp
	}
}

func generateSessionID() string {
	sessionID := uuid.New().String()
	return sessionID
}
