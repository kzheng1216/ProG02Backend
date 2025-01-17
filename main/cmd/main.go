package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// User 定义返回的用户结构体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 创建数据库连接
func getDB() (*sql.DB, error) {
	// 替换为你的 MySQL 配置信息
	dsn := "mysql:Zaq1@wsx@tcp(106.15.33.153:3306)/mydb01_dev"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// 获取用户信息
func getUser(c *gin.Context) {
	id := c.Param("id")

	// 从 MySQL 中查询用户
	db, err := getDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	var user User
	query := "SELECT id, name, email FROM user WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	// 返回查询到的用户信息
	c.JSON(http.StatusOK, user)
}

func main() {
	// 初始化 Gin 路由
	r := gin.Default()

	// 定义 API 路由
	r.GET("/api/user/:id", getUser)

	// 启动服务
	err := r.Run(":9882")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
