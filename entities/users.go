package entities

import "database/sql"

type Users struct {
	UserId   string
	Name     string
	Phone    string
	Password string
	Balance  sql.NullInt64
}
