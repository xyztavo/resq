package utils

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/configs"
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
