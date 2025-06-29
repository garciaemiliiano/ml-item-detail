package items

import (
	"item-detail-api/src/infrastructure/dbtests"
	dbtest "item-detail-api/src/infrastructure/dbtests"
	"item-detail-api/src/repositories/basemodel"
	"os"
	"testing"

	"github.com/google/uuid"
)

var suite *dbtest.Suite

func TestMain(m *testing.M) {
	suite = dbtest.NewSuite()
	suite.Start()
	suite.AddTables(ItemDAO{})
	code := m.Run()
	os.Exit(code)
}

func populateItems(ids []uuid.UUID) error {
	var err error
	for _, id := range ids {
		err = dbtests.CreateTestDaos(suite, []ItemDAO{{
			BaseModel: basemodel.BaseModel{
				ID:        id,
				CreatedAt: "2023-10-01 12:00:00",
				UpdatedAt: "2023-10-01 12:00:00",
			},
			ProductID:    uuid.New(),
			ProviderID:   uuid.New(),
			SerialNumber: "SN-" + id.String(),
			Stock:        100,
		}})
	}
	return err
}

func populate(t *testing.T, ids ...uuid.UUID) {
	err := populateItems(ids)
	if err != nil {
		t.Fatalf("failed to populate items: %v", err)
	}
}
