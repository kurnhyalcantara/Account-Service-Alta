package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func AddUser(db *sql.DB, user entities.Users) string {
	return ""
}

func LoginUser(db *sql.DB, phone string, password string) (string, error) {
	// Memeriksa apakah ada user lain yang sedang login saat ini
	if checkLoginActivity(db) {
		return "", fmt.Errorf("LoginUser: %s", "Terdapat User Lain yang sedang login")
	}

	// Memeriksa apakah phone terdaftar di database
	if !verifyPhoneRegistered(db, phone) {
		return "", fmt.Errorf("LoginUser: %s", "Nomor anda tidak terdaftar")
	}

	// Memeriksa kredensial
	credential, err := db.Query("SELECT password FROM users WHERE phone = ?", phone)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if credential.Next() {
		var passwordRegistered string
		errScan := credential.Scan(&passwordRegistered)
		if errScan != nil {
			return "", fmt.Errorf(errScan.Error())
		}
		if passwordRegistered != password {
			return "", fmt.Errorf("LoginUser: %s", "Password tidak cocok")
		}
	}

	// Store login activity
	// Generate unique id dengan menggunakan library gonanoid
	id, errNano := gonanoid.New(16)
	if errNano != nil {
		return "", fmt.Errorf("LoginUser: %v", errNano)
	}
	loginActivityId := "activityLogin-" + id
	_, errInsert := db.Exec("INSERT INTO login_activity (activity_id, phone) VALUES (?, ?)", loginActivityId, phone)
	// fmt.Println(errInsert.Error())
	if errInsert != nil {
		return "", fmt.Errorf("LoginUser: %v", errInsert)
	}
	fmt.Println("Berhasil insert")
	return phone, nil
}

func checkLoginActivity(db *sql.DB) bool {
	query, err := db.Query("SELECT * FROM login_activity")
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	if query.Next() {
		return true
	}
	return false
}

func verifyPhoneRegistered(db *sql.DB, phone string) bool {
	query, err := db.Query("SELECT phone FROM users WHERE phone = ?", phone)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}	
	
	if query.Next() {
		return true
	}
	return false
}

func updateUser(db *sql.DB, user entities.Users) entities.Users {
	return entities.Users{}
}

func DeleteUser(db *sql.DB, phone string) {
	_, err := db.Exec("DELETE FROM users WHERE phone = ?", phone)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// func SearchUser(db *sql.DB, phone string) {
// 	query, err := db.Query("SELECT name, phone FROM users WHERE phone = ?", phone)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
	
// }

