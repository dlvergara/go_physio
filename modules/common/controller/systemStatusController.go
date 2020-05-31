package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type SystemStatusResponse struct {
	Version  string `json:"version"`
	AppDate string `json:"app_date"`
}

func SystemStatus(c echo.Context) (err error) {
	response := new(SystemStatusResponse)
	response.Version = "3.0.3"
	response.AppDate = time.Now().Format("2006-01-02 15:04:05")

	return c.JSON(http.StatusOK, response)
}