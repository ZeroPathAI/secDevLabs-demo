package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// ValidateUserAccess checks if the authenticated user has access to the requested resource
func ValidateUserAccess(c echo.Context, requestedID string) error {
	authenticatedUserID := c.Get("user")
	if authenticatedUserID == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
			"result": "error", 
			"details": "Authentication required",
		})
	}

	if authenticatedUserID.(string) != requestedID {
		return echo.NewHTTPError(http.StatusForbidden, map[string]string{
			"result": "error",
			"details": "Access denied",
		})
	}

	return nil
}
