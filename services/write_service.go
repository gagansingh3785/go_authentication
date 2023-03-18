package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func WriteService(req requests.WriteRequest, username string) responses.WriteResponse {
	resp := responses.WriteResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	articleID, err := repository.CreateNewArticle(username, req.Title, req.Text)
	if err != nil || articleID == "" {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}
