package response

import "github.com/hanifbg/category-product/service/product"

type getAllProductResponse struct {
	Product []GetProductResponse `json:"product"`
}

func NewGetAllProductResponse(product []product.Product) getAllProductResponse {
	getAllProductResponse := getAllProductResponse{}

	for _, value := range product {
		var getProductResponse GetProductResponse

		getProductResponse.ID = int(value.ID)
		getProductResponse.CategoryID = int(value.CategoryID)
		getProductResponse.Name = value.Name
		getProductResponse.Price = int(value.Price)
		getProductResponse.Image = value.Image
		getProductResponse.CreatedAt = value.CreatedAt
		getProductResponse.UpdatedAt = value.UpdatedAt

		getAllProductResponse.Product = append(getAllProductResponse.Product, getProductResponse)
	}

	if getAllProductResponse.Product == nil {
		getAllProductResponse.Product = []GetProductResponse{}
	}

	return getAllProductResponse
}
