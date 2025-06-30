package seed

import (
	"item-detail-api/config"
	"item-detail-api/src/repositories/basemodel"
	"item-detail-api/src/repositories/categories"
	"item-detail-api/src/repositories/items"
	"item-detail-api/src/repositories/products"
	"item-detail-api/src/repositories/providers"
	"item-detail-api/src/repositories/reviews"
	"item-detail-api/src/repositories/users"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, logger *zap.Logger) {
	if config.GetConfig().SEED != "true" {
		logger.Info("seeding is disabled, skipping migration")
		return
	}
	migrateProducts(db, logger)
	migrateCategories(db, logger)
	migrateProviders(db, logger)
	migrateUsers(db, logger)
	migrateItems(db, logger)
	migrateReviews(db, logger)
}

func migrateProducts(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&products.ProductDAO{}); err != nil {
		logger.Error("failed to migrate Product model", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&products.ProductDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count products", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: products already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	products := []products.ProductDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("7d6db107-6e2e-4fa4-b1b5-c7fa874937f6"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Price:       349.99,
			Name:        "Auriculares Bluetooth",
			Description: "Auriculares inalámbricos con cancelación de ruido",
			CategoryID:  uuid.MustParse("1e0f49e2-df51-4122-a5b5-7c2d68dbec65"),
			ImageURL:    "https://img.freepik.com/free-photo/wireless-earbuds-with-neon-cyberpunk-style-lighting_23-2151074296.jpg?t=st=1751183790~exp=1751187390~hmac=860cea2ca3dd1ca51fd3552003a3f199a18afa4d273c0a2204f872fe174930dd&w=2000",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("e014c2d1-391e-4ff6-9503-bb27b1c2d8a4"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Price:       99.0,
			Name:        "Lámpara LED",
			Description: "Lámpara de bajo consumo con diseño moderno",
			CategoryID:  uuid.MustParse("bf6f7f79-2652-41b4-b61e-9b7ae52b8b5f"),
			ImageURL:    "https://http2.mlstatic.com/D_NQ_NP_2X_835112-MLA80947206355_112024-F.webp",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("7f505987-bb47-4a13-b1eb-e16de8e9a8ec"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Price:       159.75,
			Name:        "Set de bloques",
			Description: "Juguete educativo de construcción para niños",
			CategoryID:  uuid.MustParse("cb0e7c3a-dfd5-4c5a-91b7-180ac4ec59f3"),
			ImageURL:    "https://http2.mlstatic.com/D_NQ_NP_2X_947343-MLU79145025588_092024-F.webp",
		},
	}

	if err := db.Create(&products).Error; err != nil {
		logger.Error("failed to seed products", zap.Error(err))
		return
	}

	logger.Info("seed completed: products created", zap.Int("count", len(products)))
}

func migrateCategories(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&categories.CategoryDAO{}); err != nil {
		logger.Error("failed to migrate categories", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&categories.CategoryDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count category", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: categories already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	categories := []categories.CategoryDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("1e0f49e2-df51-4122-a5b5-7c2d68dbec65"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Electrónica",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("bf6f7f79-2652-41b4-b61e-9b7ae52b8b5f"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Hogar",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("cb0e7c3a-dfd5-4c5a-91b7-180ac4ec59f3"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Juguetes",
		},
	}

	if err := db.Create(&categories).Error; err != nil {
		logger.Error("failed to seed categories", zap.Error(err))
		return
	}

	logger.Info("seed completed: categories created", zap.Int("count", len(categories)))
}

func migrateProviders(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&providers.ProviderDAO{}); err != nil {
		logger.Error("failed to migrate providers", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&providers.ProviderDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count providers", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: providers already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	providers := []providers.ProviderDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("3e83e38c-92cb-4302-8dc0-4c5b725de8a4"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Tecno S.A.",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("a99c861f-6e29-41a1-99a6-cd3e9df2f97a"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Importadora Gómez",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("d7f6e4c0-ec8e-48a4-87f4-98475900dddc"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Name: "Juegos Mundiales",
		},
	}

	if err := db.Create(&providers).Error; err != nil {
		logger.Error("failed to seed providers", zap.Error(err))
		return
	}

	logger.Info("seed completed: providers created", zap.Int("count", len(providers)))
}

func migrateUsers(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&users.UserDAO{}); err != nil {
		logger.Error("failed to migrate users", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&users.UserDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count users", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: users already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	users := []users.UserDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("b6f8775d-4c89-4a67-8b89-9bb34c1a5c5d"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Username: "cristian",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("a3c5b8e9-2026-41a3-9c01-6c7b0385ad8f"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Username: "sofia",
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("9a1fa84a-54a5-4fbd-8d0d-25884e3f542e"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Username: "lucas",
		},
	}

	if err := db.Create(&users).Error; err != nil {
		logger.Error("failed to seed users", zap.Error(err))
		return
	}

	logger.Info("seed completed: users created", zap.Int("count", len(users)))
}

func migrateItems(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&items.ItemDAO{}); err != nil {
		logger.Error("failed to migrate items", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&items.ItemDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count items", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: items already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	items := []items.ItemDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("ec24e26d-80ea-4069-bfcb-4205aef7dd2a"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			ProductID:    uuid.MustParse("7d6db107-6e2e-4fa4-b1b5-c7fa874937f6"),
			ProviderID:   uuid.MustParse("3e83e38c-92cb-4302-8dc0-4c5b725de8a4"),
			SerialNumber: "SN-AUR-001",
			Stock:        15,
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("3c00eb10-1de4-4f3e-b18d-d5e3a1f3c00b"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			ProductID:    uuid.MustParse("e014c2d1-391e-4ff6-9503-bb27b1c2d8a4"),
			ProviderID:   uuid.MustParse("a99c861f-6e29-41a1-99a6-cd3e9df2f97a"),
			SerialNumber: "SN-LAM-003",
			Stock:        40,
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("9cb8c7b2-7bb2-44a7-9b41-97f4601212fa"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			ProductID:    uuid.MustParse("7f505987-bb47-4a13-b1eb-e16de8e9a8ec"),
			ProviderID:   uuid.MustParse("d7f6e4c0-ec8e-48a4-87f4-98475900dddc"),
			SerialNumber: "SN-JUG-005",
			Stock:        25,
		},
	}

	if err := db.Create(&items).Error; err != nil {
		logger.Error("failed to seed items", zap.Error(err))
		return
	}

	logger.Info("seed completed: items created", zap.Int("count", len(items)))
}

func migrateReviews(db *gorm.DB, logger *zap.Logger) {
	if err := db.AutoMigrate(&reviews.ReviewDAO{}); err != nil {
		logger.Error("failed to migrate items", zap.Error(err))
		return
	}

	var count int64
	if err := db.Model(&reviews.ReviewDAO{}).Count(&count).Error; err != nil {
		logger.Error("failed to count items", zap.Error(err))
		return
	}
	if count > 0 {
		logger.Info("seed skipped: items already exist")
		return
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	updatedAt := now.Format("2006-01-02 15:04:05")

	reviews := []reviews.ReviewDAO{
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("e5093cbf-d3c6-4c87-a5d1-b6d8cf313da0"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Description: "Excelente calidad de sonido",
			UserID:      uuid.MustParse("b6f8775d-4c89-4a67-8b89-9bb34c1a5c5d"),
			ProductID:   uuid.MustParse("7d6db107-6e2e-4fa4-b1b5-c7fa874937f6"),
			Rating:      3,
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("a0a5a882-cb6e-46e6-b527-46f59379b88f"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Description: "Muy útil y práctica",
			UserID:      uuid.MustParse("a3c5b8e9-2026-41a3-9c01-6c7b0385ad8f"),
			ProductID:   uuid.MustParse("e014c2d1-391e-4ff6-9503-bb27b1c2d8a4"),
			Rating:      2,
		},
		{
			BaseModel: basemodel.BaseModel{
				ID:        uuid.MustParse("9d2a04e6-cf6b-4d15-9a43-e3a741f9ed21"),
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			Description: "Perfecto para niños pequeños",
			UserID:      uuid.MustParse("9a1fa84a-54a5-4fbd-8d0d-25884e3f542e"),
			ProductID:   uuid.MustParse("7f505987-bb47-4a13-b1eb-e16de8e9a8ec"),
			Rating:      4,
		},
	}

	if err := db.Create(&reviews).Error; err != nil {
		logger.Error("failed to seed reviews", zap.Error(err))
		return
	}

	logger.Info("seed completed: reviews created", zap.Int("count", len(reviews)))
}
