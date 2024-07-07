package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"github.com/xyztavo/resq/internal/utils"
)

func CreateCompany(c echo.Context) error {
	userId, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	companyBody := new(models.CreateCompanyBody)
	if err := utils.BindAndValidate(c, companyBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdCompanyId, err := database.CreateCompany(userId, companyBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message":          "company created with ease!",
		"createdCompanyId": createdCompanyId,
	})
}

func GetCompanies(c echo.Context) error {
	companies, err := database.GetCompanies()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, companies)
}

func GetUserCompany(c echo.Context) error {
	claims, err := utils.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	company, err := database.GetUserCompany(claims.OrgId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, company)
}

func GetUserCompanyAdmin(c echo.Context) error {
	userId, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	company, err := database.GetUserAdminCompany(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, company)
}
