package services

import (
	"ProG02Backend/main/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// User 定义返回的用户结构体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 获取用户信息
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /api/user/{id}")
	// 从 URL 参数中获取用户 ID
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("id: " + id)

	// 从 MySQL 中查询用户
	db, err := utils.GetDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	query := "SELECT id, name, email FROM user WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	// 设置响应类型为 JSON
	w.Header().Set("Content-Type", "application/json")
	// 返回查询到的用户信息
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /api/users")
	db, err := utils.GetDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 查询数据库
	rows, err := db.Query("SELECT id, name, email FROM user")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to fetch data from the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 创建一个切片来存储查询到的用户
	var users []User

	// 遍历查询结果
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Fatal(err)
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// 检查遍历过程中是否发生错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		http.Error(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 将查询结果转为 JSON 并写入响应体
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Fatal(err)
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
}
