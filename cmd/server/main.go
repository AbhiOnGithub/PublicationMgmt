package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/abhiongithub/publicationmgmt/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our App")

	handler := transportHttp.NewHandler()
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
