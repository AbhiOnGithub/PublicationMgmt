package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}

func main() {
	fmt.Printf(("Publication Mgmt"))
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
