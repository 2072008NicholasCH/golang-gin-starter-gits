// Copyright 2021 The starter Authors. All rights reserved.
// This is an API Gateway Resource for starter
// Built with gRPC and Gin Gonic
//
// Auth Resource
package resource

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{Token: token}
}

type RegisterRequest struct {
	Name     string `form:"name,omitempty" json:"name,omitempty" binding:"required"`
	Email    string `form:"email,omitempty" json:"email,omitempty" binding:"required"`
	Password string `form:"password,omitempty" json:"password,omitempty" binding:"required"`
}
