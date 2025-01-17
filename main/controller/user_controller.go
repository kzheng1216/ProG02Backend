package controller

import (
	"ProG02Backend/main/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserController struct{}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("GET /api/user/{id} | id: " + id)

	user, _ := service.GetUser(id)

	// 设置响应类型为 JSON
	w.Header().Set("Content-Type", "application/json")
	// 返回查询到的用户信息
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (c UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /api/users")
	users, _ := service.GetAllUser()

	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 将查询结果转为 JSON 并写入响应体
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Fatal(err)
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
	}
}
