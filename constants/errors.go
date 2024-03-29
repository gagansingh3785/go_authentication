package constants

import "github.com/pkg/errors"

const (
	//Return string errors
	InternalServerError = "Internal Server Error"
	InvalidCredentials  = "The entered credentials are invalid. Please try again :)"
	SessionExpired      = "Sorry you session expired, please login again"
	UsernameTaken       = "This username is already taken, please try with a different username"
	EmailAlreadyTaken   = "This email is already taken, please try with a different email"
	BadRequest          = "please provide all the fields"
	ArticlePageNotFound = "The given page number is invalid"
	InvalidTags         = "The provided tags are invalid"
	InvalidArticleID    = "The given article id is invalid"

	//Messages
	ArticleNotLikedMessage = "The given articles has not been like by you"

	//error codes for postgres driver
	UniqueConstraintViolation = "23505"
)

var (
	//SQL errors
	ErrSQLNoRows                 = errors.Errorf("%s", "Error no result for query")
	ErrEmptyTags                 = errors.Errorf("%s", "No tag present in request")
	ErrExcessiveTags             = errors.Errorf("%s", "More than 3 tags supplied")
	ErrEmptyTitle                = errors.Errorf("%s", "No title provided")
	ErrEmptyArticleBody          = errors.Errorf("%s", "Empty body for article")
	ErrEmptyCommentContent       = errors.Errorf("%s", "The comment cannot be empty")
	ErrInvalidRequestParameters  = errors.Errorf("%s", "The request is invalid")
	ErrUniqueConstraintViolation = errors.Errorf("%s", "You have already like this article")
)
