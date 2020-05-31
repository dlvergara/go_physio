package router

import (
	systemStatusController "physiobot/modules/common/controller"
	"physiobot/modules/common/errors"
	common "physiobot/modules/common/middleware"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

// Init routes and auth middleware
func Init(server *echo.Echo) *echo.Echo {

	// Set custom context
	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &common.CustomContext{c}
			return h(cc)
		}
	})
	/*
	// Check authorization header
	server.Use(common.AuthHeader)

	// Basic auth middleware
	server.Use(middleware.BasicAuthWithConfig(
		middleware.BasicAuthConfig{
			Skipper:   common.SkipBasic,
			Validator: common.IsAuthBasic,
			Realm:     "Restricted"}))

	// Bearer auth middleware
	server.Use(middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			Skipper:    common.SkipBearer,
			Validator:  common.IsAuthBearer,
			KeyLookup:  "header:" + echo.HeaderAuthorization,
			AuthScheme: "Bearer"}))
	*/
	// Custom Error handler
	server.HTTPErrorHandler = errors.ErrorHandler

	// Custom request data validator
	server.Validator = common.Validator()

	// API version v3
	v3 := server.Group("/v1")

	//Routes

	// System status
	v3.GET("/system/status", systemStatusController.SystemStatus)

	/*
	// Endpoints that were initially for Walmart but can be used for other merchants
	v3.GET("/customs/duty-id/:duty_id", dutyIdController.GetDutyByID)
	v3.GET("/customs/duty-id", dutyIdController.GetDutyByProduct)
	v3.POST("/order/cancel", orderController.CancelOrder)
	v3.POST("/order/:orderID/cancel-items", orderController.CancelItems)
	v3.GET("/order/status", orderController.OrderDetail)
	v3.POST("/distribution/tracking-update", trackingController.TrackingUpdate)
	v3.DELETE("/order/:orderID/item-refund", orderController.ItemRefund)

	// Order Report
	v3.GET("/order/report", reportController.GetOrderReport)

	// Connect Landed Cost
	v3.POST("/customs/landedcost", landedCostController.LandedCost)

	// Edt Cache engine
	v3.POST("/customs/edt", edtCacheController.EdtController)
	v3.POST("/customs/edt-cache", edtCacheController.CacheController)
	v3.GET("/customs/edt-frequency-report", edtCacheController.FrequencyReportController)
	v3.GET("/customs/edt-frequency-report-per-country", edtCacheController.FrequencyReportPerCountryController)

	v3.GET("/cst/order-list", cstControllers.GetOrderList)
	 */
	/*
		// API version v3
		v1 := server.Group("/v1")

		// Rest version of World Tariff endpoints
		v1.POST("/wt/edt", wtController.ServiceEDT)
		v1.POST("/wt/edths6", wtController.ServiceEDTHS6)
		v1.POST("/wt/eccn", wtController.ServiceECCN)
		v1.POST("/wt/hssearch", wtController.ServiceHSSearch)
		v1.POST("/wt/hsdetail", wtController.ServiceHSDetailLookup)
		v1.POST("/wt/scheduleb", wtController.ServiceScheduleB)
		v1.POST("/wt/rps", wtController.ServiceRPS)
	*/

	return server
}
