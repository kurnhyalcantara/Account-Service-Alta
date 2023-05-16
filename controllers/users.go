package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"fmt"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// AddUser berfungsi untuk menambahkan data pengguna baru ke dalam database.
// Fungsi ini menerima parameter db yang merupakan objek database yang sudah terkoneksi,
// dan user yang merupakan data pengguna yang ingin ditambahkan ke database.
// Fungsi ini mengembalikan ID pengguna yang baru saja ditambahkan ke database.
func AddUser(db *sql.DB, user entities.Users) (string, error) {
	// generate unique ID untuk pengguna baru
	userId, err := gonanoid.New(16)
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}
	// lakukan query untuk menyimpan data pengguna ke dalam database
	user.UserId = userId
	_, err = db.Exec("INSERT INTO users(user_id, name, phone, password) VALUES (?, ?, ?, ?)", user.UserId, user.Name, user.Phone, user.Password)
	if err != nil {
		return "", fmt.Errorf("failed to add user to database: %v", err)
	}
	return userId, nil
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

