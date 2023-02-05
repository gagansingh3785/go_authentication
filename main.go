package main

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/config"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/server"
	_ "github.com/lib/pq" // this import is required to fetch the complete implementation of postgres driver for sql.Open()
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
)

func setupAuthentication() error {
	defer database.DBConn.Close()
	fmt.Println("Setting up authentication_app")

	//Creating DB connection
	//err := database.CreateNewDBConn()
	//if err != nil {
	//	panic(err)
	//}

	//database.TestDatabaseConnection()
	//err = database.DBConn.Ping()
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("successfully created the database connection")

	//Router: This returns a gorilla/mux router which takes paths and routes them to handlers
	r := server.CreateNewRouter()

	//Server: This is a basic http server
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", constants.AUTHENTICATION_HOST, constants.PORT),
		Handler: r,
	}

	//config: Setting up config for app
	config.InitConfig()

	fmt.Printf("%v", config.GlobalConfig.SendGrid)

	//Starting the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%v", err)
	}
	return nil
}

func main() {
	fmt.Println("This is a authentication microservice")

	//Creating CLI ap
	app := cli.NewApp()
	app.Name = "authentication_app"
	app.Version = "0.0"

	// This action is run when no subcommand is provided
	app.Action = func(*cli.Context) error {
		err := setupAuthentication()
		panic(err.Error())
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:        "start",
			Description: "Starting the http server",
			Action: func(*cli.Context) error {
				err := setupAuthentication()
				panic(err.Error())
				return nil
			},
		},
		{
			Name:        "migrations:run",
			Description: "Applying migrations to Database",
			Action: func(*cli.Context) error {
				err := database.ApplyMigrations()
				if err != nil {
					fmt.Printf("%s", err.Error())
				}
				return err
			},
		},
		{
			Name:        "migrations:create",
			Description: "Creating migration files",
			Action: func(c *cli.Context) error {
				filename := c.Args().Get(0)
				err := database.CreateMigrationFiles(filename)
				if err != nil {
					fmt.Printf("%s", err.Error())
				}
				return err
			},
		},
		{
			Name:        "migrations:rollback",
			Description: "Rolling back the last migration",
			Action: func(*cli.Context) error {
				return database.RollbackMigration()
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
