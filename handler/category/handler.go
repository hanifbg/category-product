package category

import (
	"github.com/hanifbg/category-product/common"
	"github.com/hanifbg/category-product/service/category"

	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service category.Service
}

func NewHandler(service category.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) UserHandler(c echo.Context) error {
	err := handler.service.GetCategory()
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
