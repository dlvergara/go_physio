package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"physiobot/config"
	"physiobot/modules/common/logger"
)

func ProcessData(c echo.Context) error {
	topics := map[int]string{};

	topics[1] = "CocosyLucas";
	topics[2] = "JetPeru";
	topics[3] = "TKambio";
	topics[4] = "CambistaInca";

	configuration := config.GetConfig()
	fmt.Println(configuration);
	//fmt.Println(topics);

	for index, element := range topics {
		logger.Info(index);
		logger.Info(element);
	}

	return c.JSON(http.StatusNoContent, "")
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
