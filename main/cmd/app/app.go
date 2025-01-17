package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const DefaultServerAddress = ":9882"

type App struct {
	*mux.Router
}

func NewApp() *App {
	app := new(App)
	app.initRouter()
	return app
}

func (app *App) Start() {
	// 启动 HTTP 服务器
	fmt.Println("Server is running on http://localhost:" + DefaultServerAddress)
	err := http.ListenAndServe(DefaultServerAddress, app.Router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func (app *App) Stop() {
	fmt.Println("Server is stopped.")
}
