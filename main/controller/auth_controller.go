package controller

import (
	"ProG02Backend/main/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthController struct{}

func (c AuthController) ValidateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /auth/login")
	// 解析请求体
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := service.ValidateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// 返回 JWT
	response := JwtResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
