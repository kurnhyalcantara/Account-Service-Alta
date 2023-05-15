package controllers

import (
	"alta/account-service-app/entities"
	"database/sql"
)

func addUser(db *sql.DB, user entities.Users) string {
	return ""
}

func loginUser(db *sql.DB, phone string, password string) string {
	return ""
}

func updateUser(db *sql.DB, user entities.Users) entities.Users {
	return entities.Users{}
}

func deleteUser(db *sql.DB, phone string) string {
	return ""
}

