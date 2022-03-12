package main

import (
	"fmt"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up the app")
	return nil
}

func main() {
	fmt.Println("Comments service")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting app")
		fmt.Println(err)
	}
}
