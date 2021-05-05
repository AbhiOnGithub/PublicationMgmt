package main

import (
	"fmt"
	"net/http"

	"github.com/abhiongithub/publicationmgmt/internal/database"
	"github.com/abhiongithub/publicationmgmt/internal/services"
	transportHttp "github.com/abhiongithub/publicationmgmt/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our App")

	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	bookService := services.NewService(db)

	handler := transportHttp.NewHandler(bookService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil

}

func main() {
	fmt.Println(("Publication Mgmt"))
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
