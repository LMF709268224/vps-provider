// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package model

import (
	"time"
)

type LoginLog struct {
	ID            int64     `db:"id" json:"id"`
	LoginUsername string    `db:"login_username" json:"login_username"`
	IpAddress     string    `db:"ip_address" json:"ip_address"`
	LoginLocation string    `db:"login_location" json:"login_location"`
	Browser       string    `db:"browser" json:"browser"`
	Os            string    `db:"os" json:"os"`
	Status        int32     `db:"status" json:"status"`
	Msg           string    `db:"msg" json:"msg"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type OperationLog struct {
	ID               int64     `db:"id" json:"id"`
	Title            string    `db:"title" json:"title"`
	BusinessType     int32     `db:"business_type" json:"business_type"`
	Method           string    `db:"method" json:"method"`
	RequestMethod    string    `db:"request_method" json:"request_method"`
	OperatorType     int32     `db:"operator_type" json:"operator_type"`
	OperatorUsername string    `db:"operator_username" json:"operator_username"`
	OperatorUrl      string    `db:"operator_url" json:"operator_url"`
	OperatorIp       string    `db:"operator_ip" json:"operator_ip"`
	OperatorLocation string    `db:"operator_location" json:"operator_location"`
	OperatorParam    string    `db:"operator_param" json:"operator_param"`
	JsonResult       string    `db:"json_result" json:"json_result"`
	Status           int32     `db:"status" json:"status"`
	ErrorMsg         string    `db:"error_msg" json:"error_msg"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

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
