package items

import (
	"item-detail-api/src/core/entities/categories"
	entity "item-detail-api/src/core/entities/items"
	"item-detail-api/src/core/entities/products"
	"item-detail-api/src/core/entities/providers"
	"item-detail-api/src/core/entities/reviews"
)

type GetItemResponse struct {
	ID           string           `json:"id"`
	ProductID    string           `json:"product_id"`
	ProviderID   string           `json:"provider_id"`
	SerialNumber string           `json:"serial_number"`
	Stock        int              `json:"stock"`
	CreatedAt    string           `json:"created_at"`
	UpdatedAt    string           `json:"updated_at"`
	DeletedAt    *string          `json:"deleted_at"`
	Product      ProductResponse  `json:"product"`
	Provider     ProviderResponse `json:"provider"`
}

type ProductResponse struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	Description string           `json:"description"`
	ImageURL    string           `json:"image_url"`
	CategoryID  string           `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Reviews     []ReviewResponse `json:"reviews,omitempty"`
}

type ReviewResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

type ProviderResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toProductResponse(p products.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID.String(),
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		CategoryID:  p.CategoryID.String(),
		Category:    toCategoryResponse(p.Category),
		Reviews:     toReviewsResponse(p.Reviews),
	}
}

func toReviewResponse(r reviews.Review) ReviewResponse {
	return ReviewResponse{
		ID:          r.ID.String(),
		Description: r.Description,
		Rating:      r.Rating,
	}
}

func toReviewsResponse(rs []reviews.Review) []ReviewResponse {
	reviewsResponse := make([]ReviewResponse, 0)
	for _, r := range rs {
		reviewsResponse = append(reviewsResponse, toReviewResponse(r))
	}
	return reviewsResponse
}

func toCategoryResponse(c categories.Category) CategoryResponse {
	return CategoryResponse{
		ID:   c.ID.String(),
		Name: c.Name,
	}
}

func toProviderResponse(p providers.Provider) ProviderResponse {
	return ProviderResponse{
		ID:   p.ID.String(),
		Name: p.Name,
	}
}

func ToResponse(i entity.Item) GetItemResponse {
	return GetItemResponse{
		ID:           i.ID.String(),
		ProductID:    i.ProductID.String(),
		ProviderID:   i.ProviderID.String(),
		SerialNumber: i.SerialNumber,
		Stock:        i.Stock,
		CreatedAt:    i.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    i.UpdatedAt.Format("2006-01-02 15:04:05"),
		Product:      toProductResponse(i.Product),
		Provider:     toProviderResponse(i.Provider),
	}
}
