package service

import (
	"ProG02Backend/main/utils"
	"fmt"
	"strconv"
)

func GetUser(id string) (User, error) {
	var user User
	user = User{}

	//Get User in Redis
	redisClient := NewRedisClient()
	user, err := redisClient.GetUser("user_" + id)
	if err == nil {
		fmt.Println("Get user in Redis: ", user)
		return user, err
	}

	//Get User in MySQL
	db, err := utils.GetDB()
	if err != nil {
		fmt.Println("Database connection failed. Error: ", err)
		return user, err
	}

	defer db.Close()
	query := "SELECT id, name, email FROM user WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		fmt.Println("Failed to fetch user in MySQL. Error: ", err)
		user.ID, _ = strconv.Atoi(id) // Convert string to int
		user.Name = "name_" + id
		user.Email = "email_" + id
		fmt.Println("Save user in MySQL: ", user)
		stmt, _ := db.Prepare("INSERT INTO user (id, name, email) VALUES (?, ?, ?)")
		_, err = stmt.Exec(user.ID, user.Name, user.Email)
	}

	//Save User in Redis
	err = redisClient.SetUser("user_"+id, user)
	if err != nil {
		fmt.Println("Failed to save user in Redis. Error: ", err)
		return user, err
	}
	fmt.Println("Save user in Redis: ", user)

	return user, nil
}

func GetAllUser() ([]User, error) {
	var users []User

	db, err := utils.GetDB()
	if err != nil {
		fmt.Errorf("Database connection failed" + err.Error())
		return users, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email FROM user")
	if err != nil {
		fmt.Errorf("Failed to fetch user" + err.Error())
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			fmt.Errorf("Error reading data" + err.Error())
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
