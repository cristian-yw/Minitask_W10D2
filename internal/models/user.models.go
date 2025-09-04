package models

type User struct {
	Id            int    `db:"id" json:"id"`
	Email         string `db:"email" json:"email"`
	Role          string `db:"role" json:"role"`
	Password_hash string `db:"password_hash" json:"password_hash"`
}
