package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/google/uuid"
)

func generateSessionID() string {
	sessionID := uuid.New().String()
	return sessionID
}

func isLoggedIn(userID string) (bool, error) {
	_, err := repository.GetSessionFromUserID(userID)
	switch err {
	case constants.ErrSQLNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func isDuplicateIdentifierPresent(identifier string) (bool, error) {
	_, err := repository.GetUserByEmailOrUsername(identifier)
	if err == constants.ErrSQLNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
