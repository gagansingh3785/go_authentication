package database

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
)

var DBConn *sql.DB

type DBContext struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}

func CreateNewDBConn() error {
	var err error
	dbContext := InitDBContext()
	connFormat := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbContext.host, dbContext.port, dbContext.user, dbContext.password, dbContext.dbName, dbContext.sslMode)
	DBConn, err = sql.Open("postgres", connFormat)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return err
}

func TestDatabaseConnection() error {
	rows, err := DBConn.Query("SELECT * FROM pg_roles")
	defer rows.Close()
	var name string
	var email string
	var password string
	for rows.Next() {
		err = rows.Scan(&name, &password, &email)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Println(name, email, password)
	}
	return nil
}

func InitDBContext() DBContext {
	return DBContext{
		host:     constants.DB_HOST,
		port:     constants.DB_PORT,
		user:     constants.DB_USER,
		password: constants.DB_PASSWORD,
		dbName:   constants.DB_NAME,
		sslMode:  constants.DB_SSLMODE,
	}
}
