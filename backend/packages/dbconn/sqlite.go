package dbconn

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteConnector struct {
	Filepath string
}

type SQLiteConfig struct {
	Filepath string
}

func NewSQLiteConnector(c SQLiteConfig) SQLiteConnector {
	return SQLiteConnector{
		Filepath: c.Filepath,
	}
}

func (c SQLiteConnector) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.Filepath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_, err = db.DB()
	if err != nil {
		return nil, err
	}

	return db, nil
}
