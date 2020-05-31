package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type FcbApiError struct {
	Success bool                   `json:"success"`
	Code    int                    `json:"code"`
	Status  int                    `json:"status"`
	Title   string                 `json:"title"`
	Message string                 `json:"message"`
	Detail  map[string]interface{} `json:"detail"`
}

type FcbProductError struct {
	ProductId string `json:"product_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func (he *FcbApiError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", he.Code, he.Message)
}

func NewApiError(code int, detail ...map[string]interface{}) *FcbApiError {
	apiError := &FcbApiError{Code: code, Message: GetErrorMessage(code)}
	if len(detail) > 0 {
		apiError.Detail = detail[0]
	}
	return apiError
}

func NewApiErrorWithProductErrors(code int, productErrors []FcbProductError) *FcbApiError {
	apiError := &FcbApiError{Code: code, Message: GetErrorMessage(code)}
	if len(productErrors) > 0 {
		m := map[string]interface{}{
			"product_errors": productErrors,
		}
		apiError.Detail = m
	}
	return apiError
}

func NewProductError(productId string, code int) FcbProductError {
	return FcbProductError{ProductId: productId, Code: code, Message: GetErrorMessage(code)}
}

func ErrorHandler(err error, c echo.Context) {
	var (
		code   = -1
		status = http.StatusInternalServerError
		title  = http.StatusText(status)
		msg    = "An unexpected exception occurred while processing the request"
		detail map[string]interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		detail = nil
	} else if he, ok := err.(*FcbApiError); ok {
		code = he.Code
		detail = he.Detail
	}

	// Some errors are thrown as Http errors
	code = convertHttpToFcb(code)
	if errorCodes[code] != nil {
		status = errorCodes[code]["status"].(int)
		title = http.StatusText(status)
		msg = errorCodes[code]["msg"].(string)
	}

	if code != -1 {
		_ = c.JSON(status, FcbApiError{
			Success: false,
			Code:    code,
			Status:  status,
			Title:   title,
			Message: msg,
			Detail:  detail,
		})
	}
}

func convertHttpToFcb(code int) int {
	if code == 404 {
		code = PATH_NOT_FOUND
	}
	if code == 405 {
		code = METHOD_NOT_ALLOWED
	}
	if code == 415 {
		code = UNSUPPORTED_MEDIA_TYPE
	}
	return code
}
