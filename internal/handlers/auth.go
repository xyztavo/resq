package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/configs"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c echo.Context) error {
	// get email and password from body and check password
	body := new(models.AuthUserBody)
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userFromDatabase, err := database.GetUserByEmail(body.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(body.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "password does not match")
	}
	// create a JWT with user id, role and return it
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.UserClaimsJwt{
		Id:   userFromDatabase.Id,
		Role: userFromDatabase.Role,
	})
	signedToken, err := accessToken.SignedString([]byte(configs.GetJwtSecret()))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
	})
}
