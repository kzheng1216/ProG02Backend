package app

import (
	"ProG02Backend/main/services"
	"github.com/gorilla/mux"
	_ "net/http/pprof"
)

func (app *App) initRouter() {
	r := mux.NewRouter()

	// 注册 GET 路由，匹配 `/api/user/{id}` 路径
	r.HandleFunc("/api/user/{id}", services.GetUser).Methods("GET")

	app.Router = r
}
