package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	fmt.Println("------ Echo Web Server ------")

	// Create a new Echo instance - the web server framework
	var e *echo.Echo;
	e = echo.New()

	// Define a GET route handler for "/" that returns the merchant_id from context
	e.GET("/", func(ctx echo.Context) error{
		return ctx.String(http.StatusOK, fmt.Sprintf("Welcome %s", ctx.Get("merchant_id").(string)))
	})

	// Middleware to extract 'id' query parameter and store it as 'merchant_id' in context
	e.Use(func (next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			c.Set("merchant_id", c.Request().URL.Query().Get("id"))
			return next(c)
		}
	});

	// Start the server on port 1323
	e.Start(":1323")
}