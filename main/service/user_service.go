package service

import (
	"ProG02Backend/main/utils"
	"fmt"
)

func GetUser(id string) (User, error) {
	var user User
	user = User{}
	db, err := utils.GetDB()
	if err != nil {
		fmt.Errorf("Database connection failed" + err.Error())
		return user, err
	}
	defer db.Close()

	query := "SELECT id, name, email FROM user WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		fmt.Errorf("Failed to fetch user" + err.Error())
		return user, err
	}
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
