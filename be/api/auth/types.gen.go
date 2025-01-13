// Package auth provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package auth

import (
	"time"
)

// AuthToken defines model for AuthToken.
type AuthToken struct {
	ExpiresAt time.Time `json:"expiresAt"`
	Token     string    `json:"token"`
}

// PostLogin defines model for PostLogin.
type PostLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ErrorMsg defines model for ErrorMsg.
type ErrorMsg struct {
	Error string `json:"error"`
}

// PostAdminLoginJSONRequestBody defines body for PostAdminLogin for application/json ContentType.
type PostAdminLoginJSONRequestBody = PostLogin
