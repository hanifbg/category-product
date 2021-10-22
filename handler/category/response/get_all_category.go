package response

import "github.com/hanifbg/category-product/service/category"

type getAllCategoryResponse struct {
	Category []GetCategoryResponse `json:"category"`
}

func NewGetAllCategoryResponse(category []category.Category) getAllCategoryResponse {
	getAllCategoryResponse := getAllCategoryResponse{}

	for _, value := range category {
		var getCategoryResponse GetCategoryResponse

		getCategoryResponse.ID = int(value.ID)
		getCategoryResponse.Name = value.Name
		getCategoryResponse.CreatedAt = value.CreatedAt
		getCategoryResponse.UpdatedAt = value.UpdatedAt

		getAllCategoryResponse.Category = append(getAllCategoryResponse.Category, getCategoryResponse)
	}

	if getAllCategoryResponse.Category == nil {
		getAllCategoryResponse.Category = []GetCategoryResponse{}
	}

	return getAllCategoryResponse
}
