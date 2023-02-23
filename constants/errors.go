package constants

import "github.com/pkg/errors"

const (
	//Return errors
	InternalServerError = "Internal Server Error"
	InvalidCredentials  = "The entered credentials are invalid. Please try again :)"
	SessionExpired      = "Sorry you session expired, please login again"
	AlreadyLoggedIn     = "You are already logged in"
	UsernameTaken       = "This username is already taken, please try with a different username"
	EmailAlreadyTaken   = "This email is already taken, please try with a different email"
)

var (
	//SQL errors
	ErrSQLNoRows = errors.Errorf("%s", "Error no result for query")
)
