package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fatih/color"
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
	user.UserId = "user-" + userId
	_, err = db.Exec("INSERT INTO users(user_id, name, phone, password) VALUES (?, ?, ?, ?)", user.UserId, user.Name, user.Phone, user.Password)
	if err != nil {
		return "", fmt.Errorf("failed to add user to database: %v", err)
	}
	return user.UserId, nil
}

func LoginUser(db *sql.DB, phone string, password string) ([]string, error) {
	// Memeriksa apakah phone terdaftar di database
	if !verifyPhoneRegistered(db, phone) {
		return []string{}, fmt.Errorf("LoginUser: %s", "Nomor anda tidak terdaftar")
	}

	// Memeriksa kredensial
	var passwordRegistered, userId, name string
	err := db.QueryRow("SELECT user_id, name, password FROM users WHERE phone = ?", phone).Scan(&userId, &name, &passwordRegistered)
	if err != nil {
		return []string{}, fmt.Errorf(err.Error())
	}
	if passwordRegistered != password {
		return []string{}, fmt.Errorf("LoginUser: %s", "Kredensial tidak valid")
	}

	// Store login activity
	_, errInsert := db.Exec("INSERT INTO login_activity (user_id) VALUES (?)", userId)
	if errInsert != nil {
		return []string{}, fmt.Errorf("LoginUser: %v", errInsert)
	}

	// Get login time
	var loginAt string
	err = db.QueryRow("SELECT login_at FROM login_activity WHERE user_id = ?", userId).Scan(&loginAt)
	if err != nil {
		return []string{}, fmt.Errorf(err.Error())
	}

	return []string{name, loginAt}, nil
}

func CheckLoginSession(db *sql.DB) string {
	query, err := db.Query("SELECT user_id FROM login_activity")
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	var userId string
	if query.Next() {
		query.Scan(&userId)
	}
	return userId
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

// ShowUser menampilkan profil pengguna yang login
func ShowUser(user *entities.Users) {
	fmt.Println("Profil Pengguna:")
	fmt.Printf("Nama: %s\n", user.Name)
	fmt.Printf("Nomor Telepon: %s\n", user.Phone)
}

// GetLoggedInUser mengembalikan data pengguna berdasarkan loggedInUserID
func GetLoggedInUser(db *sql.DB, loggedInUserID string) (*entities.Users, error) {
	// Query ke database untuk mendapatkan data pengguna berdasarkan loggedInUserID
	query := "SELECT name, phone FROM users WHERE user_id = ?"
	row := db.QueryRow(query, loggedInUserID)

	var user entities.Users
	err := row.Scan(&user.Name, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Data pengguna tidak ditemukan")
		}
		return nil, err
	}

	return &user, nil
}

func updateUser(db *sql.DB, user entities.Users) entities.Users {
	return entities.Users{}
}

func DeleteUser(db *sql.DB) {
	userId := CheckLoginSession(db)
	_, err := db.Exec("DELETE FROM users WHERE user_id = ?", userId)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		color.HiGreen("Akun berhasil dihapus")
	}
}

func SearchUser(db *sql.DB, phone string) entities.Users {
	var user entities.Users
	err := db.QueryRow("SELECT name, phone FROM users WHERE phone = ?", phone).Scan(&user.Name, &user.Phone)
	if err != nil {
		log.Fatal(err.Error())
	}
	return user
}



