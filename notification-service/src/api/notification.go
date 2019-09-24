package api

import (
	errs "cinemas-microservices/notification-service/src/errors"
	"cinemas-microservices/notification-service/src/models"
	"cinemas-microservices/notification-service/src/smtp"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// SendEmail ...
func (a API) SendEmail(c echo.Context) error {
	c.Request().Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)

	n := new(models.Notification)

	if err := c.Bind(n); err != nil {
		return errs.Send("User", "Could not get Notification data", err)
	}

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{n.User.Email}

	Subject := "Do Not Reply, Cinemas Company 👥 <no-replay@cinemas.com>"
	tickets := fmt.Sprintf("<h1>Tickest for %s</h1>", n.Movie.Title)
	cinema := fmt.Sprintf("<p>Cinema: %s</p>", n.Cinema.Name)
	room := fmt.Sprintf("<p>Room: %s</p>", n.Cinema.Room)
	seat := fmt.Sprintf("<p>Seats: %s</p>", n.Cinema.Seats)
	description := fmt.Sprintf("<p>description: %s</p>", n.Description)
	total := fmt.Sprintf("<p>Total: %d</p>", n.TotalAmount)
	orderID := fmt.Sprintf("<p>Order ID: %s</p>", n.OrderID)

	message := fmt.Sprintf(`
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>
	%s

	%s
	%s
	%s

	%s

	%s
	%s

	<h3>Cinemas Microserivce 2019, Enjoy your movie !</h3>
	</body>
	</html>
	`, tickets, cinema, room, seat, description, total, orderID)

	bodyMessage := a.SMTP.WriteHTMLEmail(Receiver, Subject, message)

	if err := a.SMTP.SendMail(Receiver, Subject, bodyMessage); err != nil {
		return errs.Send("External", "An error occured sending an email", err)
	}

	res := map[string]interface{}{
		"msg": "Mail sent successfully to " + n.User.Email,
	}

	return c.JSON(http.StatusCreated, res)
}

// SendSMS ...
func (a API) SendSMS(c echo.Context) error {

	res := map[string]interface{}{
		"msg": "SMS Sent",
	}

	return c.JSON(http.StatusOK, res)
}

type (
	// API ...
	API struct {
		SMTP *smtp.Config
	}

	// Repository ...
	Repository interface {
		SendEmail(c echo.Context) error
		SendSMS(c echo.Context) error
	}
)

// Connect ...
func Connect(smtp *smtp.Config) (Repository, error) {
	api := new(API)
	api.SMTP = smtp

	return api, nil
}