package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewSvcRouter(r *mux.Router) *mux.Router {
	// 定义一个中间件函数来设置CORS响应头
	corsMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 设置允许所有来源访问，实际生产中可根据需要指定特定来源
			w.Header().Set("Access-Control-Allow-Origin", "*")
			// 设置允许的请求方法
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// 设置允许的请求头
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

			// 如果是OPTIONS预检请求，直接返回状态码200
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			// 调用下一个处理器
			next.ServeHTTP(w, r)
		})
	}

	authController := AuthController{}
	r.HandleFunc("/auth/login", corsMiddleware(authController.ValidateUser)).Methods(http.MethodPost, http.MethodOptions)

	userController := UserController{}
	r.HandleFunc("/api/user/{id}", corsMiddleware(userController.GetUser)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/users", corsMiddleware(userController.GetAllUser)).Methods(http.MethodGet, http.MethodOptions)
	return r
}
