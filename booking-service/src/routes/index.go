package routes

import (
	"cinemas-microservices/booking-service/src/api"

	"github.com/labstack/echo"
)

// API ...
func API(app *echo.Group, repo api.Repository) {
	PaymentAPI(app, repo)
}
