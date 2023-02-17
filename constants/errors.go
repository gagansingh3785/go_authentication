package constants

import "github.com/pkg/errors"

const (
	//Return errors
	InternalServerError = "Internal Server Error"
	InvalidCredentials  = "The entered credentials are invalid. Please try again :)"
)

var (
	//SQL errors
	ErrSQLNoRows = errors.Errorf("%s", "Error no result for query")
)
