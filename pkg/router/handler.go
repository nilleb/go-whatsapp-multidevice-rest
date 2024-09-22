package router

import (
	"fmt"
	"runtime/debug"

	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/log"
	"github.com/labstack/echo/v4"
)

func HttpErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(500, err.Error())
	}

	stack := string(debug.Stack())
	log.Print(c).Error(stack)

	response := &ResError{
		Status: false,
		Code:   report.Code,
		Error:  fmt.Sprintf("%v", report.Message),
	}

	logError(c, response.Code, response.Error)
	c.JSON(response.Code, response)
}
