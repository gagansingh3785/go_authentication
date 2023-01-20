package database

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"time"
)

func GetDBConnectURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", constants.DB_USER, constants.DB_PASSWORD, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME)
}

func ApplyMigrations() error {
	var migrationFilePath string
	migrationFilePath = constants.MigrationFilePath
	migrations, err := migrate.New("file://"+migrationFilePath, GetDBConnectURL())
	if err != nil {
		panic("failed to create migrations" + err.Error())
	}
	err = migrations.Up()
	return err
}

func RollbackMigration() error {
	var migrationFilePath string
	migrationFilePath = constants.MigrationFilePath
	migrations, err := migrate.New("file://"+migrationFilePath, GetDBConnectURL())
	if err != nil {
		panic("failed to create migrations " + err.Error())
	}
	err = migrations.Down()
	return err
}

func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return fmt.Errorf("emply file name given")
	}
	migrationFileVersion := time.Now().Unix()
	upFilePath := fmt.Sprintf("%s/%d_%s.up.sql", constants.MigrationFilePath, migrationFileVersion, filename)
	downFilePath := fmt.Sprintf("%s/%d_%s.down.sql", constants.MigrationFilePath, migrationFileVersion, filename)

	//creating up migration file
	up, err := os.Create(upFilePath)
	if err != nil {
		return err
	}
	fmt.Println("Up file created successfully")

	//creating down migration file
	down, err := os.Create(downFilePath)
	if err != nil {
		os.Remove(upFilePath)
		return err
	}
	fmt.Println("down file created successfully")

	defer up.Close()
	defer down.Close()

	return nil
}
