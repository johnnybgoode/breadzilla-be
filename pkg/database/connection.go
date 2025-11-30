package database

import (
	"database/sql"
)

func Connect(c *Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.GetMysql().FormatDSN())

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
