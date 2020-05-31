package common

import "github.com/labstack/echo"

// CustomContext model
type CustomContext struct {
	echo.Context
}

// GetPartnerID from context
func (c *CustomContext) GetPartnerID() int {
	return c.Get("partnerID").(int)
}
