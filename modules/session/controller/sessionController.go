package controller

import (
	"physiobot/modules/session/rest"
	"github.com/nu7hatch/gouuid"
	"github.com/labstack/echo"
	"net/http"
)

func CreateSession(c echo.Context) error {
	session := new(rest.NewSession)

	identification, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	session.SessionID = identification.String()

	return c.JSON(http.StatusOK, session)
}