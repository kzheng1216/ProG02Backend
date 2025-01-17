package main

import (
	"ProG02Backend/main/cmd/app"
)

func main() {
	application := app.NewApp()
	defer application.Stop()
	application.Start()
}
