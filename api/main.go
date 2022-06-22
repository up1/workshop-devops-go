package main

import (
	"demo/db"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	// addressDb := "postgres://data:password@db:5432/data?sslmode=disable"
	addressDb := os.Getenv("DB_URL")
	dbClient := db.NewPostgresClient(addressDb)
	h := &Handler{dbClient: dbClient}

	e := echo.New()
	e.GET("/accounts/:id", h.Home)
	e.Logger.Fatal(e.Start(":1323"))
}

type Handler struct {
	dbClient db.DbClient
}

func (h *Handler) Home(c echo.Context) error {
	accountID := c.Param("id")
	account, err := h.dbClient.InquiryAccount(c.Request().Context(), accountID)
	if err != nil {
		logrus.Errorf("Error reading accountID '%v' from DB: %v", accountID, err.Error())
		if err.Error() != "" {
			return c.JSON(http.StatusInternalServerError, []byte(err.Error()))
		} else {
			return c.JSON(http.StatusNotFound, []byte("Account not found"))
		}
	}
	return c.JSON(http.StatusOK, account)
}
