package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	OrgType  *string `json:"orgType"`
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

type CompanyAdmin struct {
	UserId    string `json:"userId"`
	CompanyId string `json:"companyId"`
}

type CreateCompanyBody struct {
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
