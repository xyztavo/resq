package handlers

import (
	"encoding/json"
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
	if err := json.NewDecoder(c.Request().Body).Decode(&userBody); err != nil {
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

func UpdateUserCompanyAdmin(c echo.Context) error {
	id, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if user.Role == "ngo_admin" {
		return echo.NewHTTPError(http.StatusBadRequest, "user already is a ngo admin. mind creating a new account for each purpose.")
	}
	if err := database.UpdateCompanyUserRole(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user is now a company admin"})
}

func UpdateUserNGOAdmin(c echo.Context) error {
	id, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if user.Role == "company_admin" {
		return echo.NewHTTPError(http.StatusBadRequest, "user already is a company admin. mind creating a new account for each purpose.")
	}
	if err := database.UpdateNGOUserRole(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user is now a NGO admin"})
}
