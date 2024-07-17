package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/configs"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/models"
	"github.com/xyztavo/resq/internal/utils"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get auth header
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "no auth header set")
		}
		authType := strings.Split(header, " ")[0]
		if authType == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "no auth type set. expected Bearer {token}")
		}
		if authType != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "auth type not valid. expected Bearer {token}")
		}
		// get token from header
		tokenFromHeader := strings.Split(header, " ")[1]
		// parse token
		claims := models.UserClaimsJwt{}
		parsedToken, err := jwt.ParseWithClaims(tokenFromHeader, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJwtSecret()), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "jwt not valid")
		}
		if !parsedToken.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "jwt not valid")
		}
		return next(c)
	}
}

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := utils.GetClaimsFromToken(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "must be admin to use the current route")
		}
		return next(c)
	}
}

func CompanyAdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := utils.GetIdFromToken(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		userFromDb, err := database.GetUserById(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if userFromDb.Role != "company_admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "must be company admin to use the current route")
		}
		return next(c)
	}
}
