package services

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func WriteService(req requests.WriteRequest, username string) responses.WriteResponse {
	resp := responses.NewWriteResponse()
	articleTags := req.Tags
	tags, err := repository.GetTagsFromName(articleTags)
	if err != nil {
		fmt.Printf("%s \n", err.Error())
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	if len(tags) == 0 {
		resp.Error = constants.InvalidTags
		resp.Message = constants.InvalidTags
		return resp
	}
	fmt.Println("tag_ids: ", tags)
	articleID, err := repository.CreateNewArticle(username, req.Title, req.Text, tags)
	if err != nil || articleID == "" {
		fmt.Printf("%s \n", err.Error())
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	err = repository.MapTagIDsWithArticleID(articleID, tags)
	if err != nil {
		fmt.Printf("%s \n", err.Error())
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}
