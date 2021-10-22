package product

import (
	"strconv"

	"github.com/hanifbg/category-product/common"
	"github.com/hanifbg/category-product/handler/product/response"
	"github.com/hanifbg/category-product/service/product"

	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service product.Service
}

func NewHandler(service product.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) ProductHandler(c echo.Context) error {
	categoryId, _ := strconv.Atoi(c.Param("category_id"))
	query := c.QueryParam("q")

	product, err := handler.service.GetProduct(categoryId, query)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllProductResponse(product)

	return c.JSON(common.NewSuccessResponse(response))
}
