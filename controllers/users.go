package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func AddUser(db *sql.DB, user entities.Users) string {
	return ""
}

func LoginUser(db *sql.DB, phone string, password string) (int64, error) {
	// Memeriksa apakah ada user lain yang sedang login saat ini
	checkUser, err := db.Query("SELECT * FROM login_activity")
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	} 

	checked := checkUser.Next()
	if checked {
		return 0, fmt.Errorf("Error")
	}

	// Memeriksa apakah phone terdaftar di database
	checkPhone, err := db.Query("SELECT phone FROM users WHERE phone = ?", phone)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	isPhoneRegistered := checkPhone != nil
	if !isPhoneRegistered {
		return 0, fmt.Errorf("Phone not registered")
	}
	
	// Generate unique id dengan menggunakan library gonanoid
	id, err := gonanoid.New(16)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	// Memeriksa kredensial
	credential := db.QueryRow("SELECT phone, password FROM users WHERE phone = ? AND password = ?", phone, password)
	var phoneRegistered, passwordRegistered string
	err = credential.Scan(&phoneRegistered, &passwordRegistered)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	if phoneRegistered != phone {
		return 0, fmt.Errorf("Akun tidak terdaftar")
	} else if passwordRegistered != password {
		return 0, fmt.Errorf("Password tidak cocok")
	}

	// Store login activity
	loginActivityId := "activityLogin-" + id
	result, err := db.Exec("INSERT INTO login_activity (activity_id, phone) VALUES (?, ?)", loginActivityId, phone)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	loginId, err := result.LastInsertId() 
	if err != nil {
		return 0, fmt.Errorf("loginUser: %v", err.Error())
	}

	return loginId, nil
}

func updateUser(db *sql.DB, user entities.Users) entities.Users {
	return entities.Users{}
}

func deleteUser(db *sql.DB, phone string) string {
	return ""
}

