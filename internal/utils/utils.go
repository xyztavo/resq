package utils

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/configs"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
)

func GetIdFromToken(c echo.Context) (userId string, err error) {
	// get auth header
	header := c.Request().Header.Get("Authorization")
	// get token from header
	tokenFromHeader := strings.Split(header, " ")[1]
	// parse token
	claims := models.UserClaimsJwt{}
	parsedToken, err := jwt.ParseWithClaims(tokenFromHeader, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJwtSecret()), nil
	})
	if err != nil {
		return "", errors.New("unauthorized from parsing jwt")
	}
	return parsedToken.Claims.(*models.UserClaimsJwt).Id, nil
}

func GetClaimsFromToken(c echo.Context) (claims models.UserClaimsJwt, err error) {
	// get auth header
	header := c.Request().Header.Get("Authorization")
	// get token from header
	tokenFromHeader := strings.Split(header, " ")[1]
	// parse token
	parsedToken, err := jwt.ParseWithClaims(tokenFromHeader, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJwtSecret()), nil
	})
	if err != nil {
		return claims, errors.New("unauthorized from parsing jwt")
	}
	if !parsedToken.Valid {
		return claims, errors.New("jwt not valid")
	}

	return claims, nil
}

func BindAndValidate(c echo.Context, structs any) error {
	if err := json.NewDecoder(c.Request().Body).Decode(&structs); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(structs); err != nil {
		return err
	}
	return nil
}

func GetUserCompanyId(c echo.Context) (companyId string, err error) {
	id, err := GetIdFromToken(c)
	if err != nil {
		return "", err
	}
	company, err := database.GetUserCompany(id)
	if err != nil {
		return "", err
	}
	return company.Id, nil
}
