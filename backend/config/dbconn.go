package config

import "item-detail-api/packages/dbconn"

func SQLiteConfig(c *Config) dbconn.SQLiteConfig {
	return dbconn.SQLiteConfig{
		Filepath: c.SQLITE_PATH,
	}
}
