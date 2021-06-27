package main

import (
	"net/http"
	"fmt"
	"os" 
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		name, _ := os.Hostname()
		result := fmt.Sprintf("Hello with host : %v", name)
		return c.String(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}