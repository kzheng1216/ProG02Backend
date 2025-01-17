package app

import (
	"ProG02Backend/main/controller"
	"github.com/gorilla/mux"
	_ "net/http/pprof"
)

func (app *App) initRouter() {
	r := mux.NewRouter()
	app.Router = controller.NewSvcRouter(r)
}
