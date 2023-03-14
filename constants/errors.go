package constants

import "github.com/pkg/errors"

const (
	//Return errors
	InternalServerError = "Internal Server Error"
	InvalidCredentials  = "The entered credentials are invalid. Please try again :)"
	SessionExpired      = "Sorry you session expired, please login again"
	UsernameTaken       = "This username is already taken, please try with a different username"
	EmailAlreadyTaken   = "This email is already taken, please try with a different email"
	BadRequest          = "Please provide all the fields"
	ArticlePageNotFound = "The given page number is invalid"
)

var (
	//SQL errors
	ErrSQLNoRows = errors.Errorf("%s", "Error no result for query")
)
