package database

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/config"
)

var DBConn *sql.DB

func CreateNewDBConn() error {
	var err error
	connFormat := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GlobalConfig.DBConfig.Host,
		config.GlobalConfig.DBConfig.Port,
		config.GlobalConfig.DBConfig.DBUser,
		config.GlobalConfig.DBConfig.DBPassword,
		config.GlobalConfig.DBConfig.DBName,
		config.GlobalConfig.DBConfig.SSLMode,
	)
	DBConn, err = sql.Open("postgres", connFormat)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return err
}
