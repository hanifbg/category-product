package handler

import (
	"github.com/hanifbg/category-product/handler/category"
	"github.com/hanifbg/category-product/handler/product"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, categoryHandler *category.Handler, productHandler *product.Handler) {

	userV1 := e.Group("v1")
	userV1.GET("/get_category", categoryHandler.CategoryHandler)
	userV1.GET("/get_product/:category_id", productHandler.ProductHandler)
}
