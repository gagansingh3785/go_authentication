package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/responses"
)

func LogoutService(sessionKey string) responses.LogoutResponse {
	resp := responses.LogoutResponse{}

	err := repository.DeleteSessionBySessionID(sessionKey)
	if err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	resp.Message = constants.LogoutSuccessful
	resp.AddAllHeaders()
	return resp
}
