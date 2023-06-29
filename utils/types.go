package utils

import (
	"time"
)

type User struct {
	ID         int64     `db:"id" json:"id"`
	Uuid       string    `db:"uuid" json:"uuid"`
	Avatar     string    `db:"avatar" json:"avatar"`
	Username   string    `db:"username" json:"username"`
	VerifyCode string    `db:"verify_code" json:"verify_code"`
	PassHash   string    `db:"pass_hash" json:"pass_hash"`
	UserEmail  string    `db:"user_email" json:"user_email"`
	Address    string    `db:"address" json:"address"`
	Role       int32     `db:"role" json:"role"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`
}
