package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserClaimsJwt struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type CreateUserBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthUserBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
