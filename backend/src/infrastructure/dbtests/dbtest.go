package dbtests

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Suite struct {
	DB     *gorm.DB
	tables []interface{}
}

func NewSuite() *Suite {
	return &Suite{}
}

func (s *Suite) Start() {
	var err error
	s.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to in-memory SQLite DB: %v", err)
	}
}

func (s *Suite) AddTables(tables ...interface{}) {
	s.tables = append(s.tables, tables...)
	if err := s.DB.AutoMigrate(tables...); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}

func (s *Suite) Clear(tables ...interface{}) {
	if len(tables) == 0 {
		tables = s.tables
	}

	tx := s.DB.Begin()
	for _, table := range tables {
		if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error; err != nil {
			tx.Rollback()
			log.Fatalf("failed to clear table: %v", err)
		}
	}
	tx.Commit()
}

func CreateTestDaos[T any](suite *Suite, daos ...T) error {
	for _, t := range daos {
		if err := suite.DB.Create(&t).Error; err != nil {
			return err
		}
	}
	return nil
}
