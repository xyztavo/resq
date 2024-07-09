package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"github.com/xyztavo/resq/internal/utils"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	userBody := new(models.CreateUserBody)
	if err := utils.BindAndValidate(c, userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdUserId, err := database.CreateUser(userBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message":       "user created with ease!",
		"createdUserId": createdUserId,
	})

}

func GetUser(c echo.Context) error {
	userId, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	user, err := database.GetUserById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
