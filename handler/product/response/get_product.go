package response

import "time"

type GetProductResponse struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
