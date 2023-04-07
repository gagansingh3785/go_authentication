package cronjobs

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/services"
	"time"
)

func ClearOldSessionJob() {
	fmt.Println("Executing Clear Session Cron Job")
	err := services.ClearOldSessionService()
	if err != nil {
		fmt.Println("Clear Session Cron Job Failure: ", err.Error())
	}
	time.Sleep(20000 * time.Second)
	ClearOldSessionJob()
}
