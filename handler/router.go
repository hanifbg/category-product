package handler

import (
	"github.com/hanifbg/category-product/handler/category"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, categoryHandler *category.Handler) {

	userV1 := e.Group("v1")
	userV1.GET("/get_category", categoryHandler.UserHandler)
}
