package handlers

import (
	"fmt"
	"net/http"

	"github.com/globocom/secDevLabs/owasp-top10-2021-apps/a1/ecommerce-api/app/db"
	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING\n")
}

// GetTicket returns the authenticated user's ticket.
func GetTicket(c echo.Context) error {
	// Get authenticated user's ID from the session/token
	authenticatedUserID := c.Get("user")
	if authenticatedUserID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"result": "error", "details": "Authentication required"})
	}

	// Get requested user ID from path parameter
	requestedID := c.Param("id")
	
	// Verify user is accessing their own ticket
	if authenticatedUserID.(string) != requestedID {
		return c.JSON(http.StatusForbidden, map[string]string{"result": "error", "details": "Access denied"})
	}

	userDataQuery := map[string]interface{}{"userID": requestedID}
	userDataResult, err := db.GetUserData(userDataQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Error finding this UserID."})
	}

	format := c.QueryParam("format")
	if format == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"result":   "success",
			"username": userDataResult.Username,
			"ticket":   userDataResult.Ticket,
		})
	}

	msgTicket := fmt.Sprintf("Hey, %s! This is your ticket: %s\n", userDataResult.Username, userDataResult.Ticket)
	return c.String(http.StatusOK, msgTicket)
}
