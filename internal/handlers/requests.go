package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"github.com/xyztavo/resq/internal/utils"
)

func CreateRequest(c echo.Context) error {
	requestBody := new(models.CreateRequest)
	if err := utils.BindAndValidate(c, requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdRequestId, err := database.CreateRequest(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"createdRequestId": createdRequestId})
}

func AcceptRequest(c echo.Context) error {
	acceptRequestBody := new(models.AcceptRequest)
	if err := utils.BindAndValidate(c, acceptRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := database.AcceptRequest(acceptRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "request updated with ease."})
}

func GetRequests(c echo.Context) error {
	requests, err := database.GetRequests()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, requests)
}
