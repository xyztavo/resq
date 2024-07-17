package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	OrgId    *string `json:"orgId"`
	Password string  `json:"password"`
}

type Company struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Rating      *float32 `json:"rating"`
	CreatorId   string   `json:"creatorId"`
}

type NGO struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Rating      *float32 `json:"rating"`
	CreatorId   string   `json:"creatorId"`
}

type Material struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsActive    bool      `json:"isActive"`
	CompanyId   string    `json:"companyId"`
}

type CreateMaterial struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CompanyAdmin struct {
	UserId    string `json:"userId"`
	CompanyId string `json:"companyId"`
}

type NGOAdmin struct {
	UserId string `json:"userId"`
	NGOId  string `json:"NGOId"`
}

type CreateCompanyBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateNGOBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UserClaimsJwt struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type CreateUserBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthUserBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
