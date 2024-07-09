package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"github.com/xyztavo/resq/internal/utils"
)

func CreateNGO(c echo.Context) error {
	userId, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	NGOBody := new(models.CreateNGOBody)
	if err := utils.BindAndValidate(c, NGOBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdNGOId, err := database.CreateNGO(userId, NGOBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message":          "NGO created with ease!",
		"createdCompanyId": createdNGOId,
	})
}

func GetUserNGO(c echo.Context) error {
	id, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userFromDb, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if userFromDb.Role != "ngo_admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "user must be ngo_admin admin")
	}
	ngo, err := database.GetUserNGO(userFromDb.OrgId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ngo)
}

func GetNGOs(c echo.Context) error {
	NGOs, err := database.GetNGOs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, NGOs)
}

func GetNGOsAdmins(c echo.Context) error {
	NGOsAdmins, err := database.GetNGOsAdmins()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, NGOsAdmins)
}
