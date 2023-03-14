package services

import (
	"github.com/gagansingh3785/go_authentication/repository"
	"time"
)

func ClearOldSessionService() error {
	currentTime := time.Now()
	err := repository.ClearOldSessions(currentTime)
	return err
}
