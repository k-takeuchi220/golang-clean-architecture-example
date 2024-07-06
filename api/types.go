// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

// UpdateUserNameRequest defines model for UpdateUserNameRequest.
type UpdateUserNameRequest struct {
	// Name 更新後の名前
	Name string `json:"name"`
}

// UpdateUserNameResponse defines model for UpdateUserNameResponse.
type UpdateUserNameResponse struct {
	// ID ユーザID
	ID string `json:"id"`

	// Name 更新後の名前
	Name string `json:"name"`
}

// UpdateUserNameJSONRequestBody defines body for UpdateUserName for application/json ContentType.
type UpdateUserNameJSONRequestBody = UpdateUserNameRequest
