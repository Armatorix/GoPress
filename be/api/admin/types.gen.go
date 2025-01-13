// Package admin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package admin

// Defines values for UserType.
const (
	ADMIN          UserType = "ADMIN"
	CONTENTCREATOR UserType = "CONTENT_CREATOR"
)

// ErrorMsg defines model for ErrorMsg.
type ErrorMsg struct {
	Error string `json:"error"`
}

// PatchUser defines model for PatchUser.
type PatchUser struct {
	Email    string   `json:"email"`
	Nick     string   `json:"nick"`
	UserType UserType `json:"userType"`
}

// PostUsers defines model for PostUsers.
type PostUsers struct {
	Email    string   `json:"email"`
	Nick     string   `json:"nick"`
	UserType UserType `json:"userType"`
}

// User defines model for User.
type User struct {
	Email    string  `json:"email"`
	Id       string  `json:"id"`
	Nick     *string `json:"nick,omitempty"`
	UserType string  `json:"userType"`
}

// UserType defines model for UserType.
type UserType string

// Users defines model for Users.
type Users = []User

// UserId defines model for userId.
type UserId = string

// UserIds defines model for userIds.
type UserIds = []string

// GetUsers defines model for GetUsers.
type GetUsers = Users

// DeleteUsersParams defines parameters for DeleteUsers.
type DeleteUsersParams struct {
	UserIds UserIds `form:"userIds" json:"userIds"`
}

// PostAdminUsersJSONRequestBody defines body for PostAdminUsers for application/json ContentType.
type PostAdminUsersJSONRequestBody = PostUsers

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = PatchUser
