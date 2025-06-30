package items

import (
	"item-detail-api/src/infrastructure/dbtests"
	dbtest "item-detail-api/src/infrastructure/dbtests"
	"item-detail-api/src/repositories/basemodel"
	"item-detail-api/src/repositories/categories"
	"item-detail-api/src/repositories/products"
	"item-detail-api/src/repositories/providers"
	"item-detail-api/src/repositories/reviews"
	"os"
	"testing"

	"github.com/google/uuid"
)

var suite *dbtest.Suite

func TestMain(m *testing.M) {
	suite = dbtest.NewSuite()
	suite.Start()
	suite.AddTables(ItemDAO{}, categories.CategoryDAO{}, products.ProductDAO{}, providers.ProviderDAO{}, reviews.ReviewDAO{})
	code := m.Run()
	os.Exit(code)
}

func populateItems(ids []uuid.UUID, productID uuid.UUID) error {
	var err error
	for _, id := range ids {
		err = dbtests.CreateTestDaos(suite, []ItemDAO{{
			BaseModel: basemodel.BaseModel{
				ID:        id,
				CreatedAt: "2023-10-01 12:00:00",
				UpdatedAt: "2023-10-01 12:00:00",
			},
			ProductID:    productID,
			ProviderID:   uuid.New(),
			SerialNumber: "SN-" + id.String(),
			Stock:        100,
		}})
	}
	return err
}

func popualateCategories(id uuid.UUID) error {
	return dbtests.CreateTestDaos(suite, []categories.CategoryDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        id,
				CreatedAt: "2023-10-01 12:00:00",
				UpdatedAt: "2023-10-01 12:00:00",
			},
			Name: "Hogar",
		},
	})
}

func populateProducts(category uuid.UUID) error {
	return dbtests.CreateTestDaos(suite, []products.ProductDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.New(),
				CreatedAt: "2023-10-01 12:00:00",
				UpdatedAt: "2023-10-01 12:00:00",
			},
			Name:        "Producto Hogar",
			Description: "Un producto para el hogar",
			CategoryID:  category,
		},
	})
}

func populate(t *testing.T, ids ...uuid.UUID) {
	categoryID := uuid.New()
	productID := uuid.New()
	err := populateItems(ids, productID)
	if err != nil {
		t.Fatalf("failed to populate items: %v", err)
	}
	err = popualateCategories(categoryID)
	if err != nil {
		t.Fatalf("failed to populate categories: %v", err)
	}
	err = populateProducts(categoryID)
	if err != nil {
		t.Fatalf("failed to populate products: %v", err)
	}
}
