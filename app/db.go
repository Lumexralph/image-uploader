package app

import (
	"database/sql"
	// register driver needed for postgreSQL
	_ "github.com/lib/pq"
)

// databaseHandler creates interface for all db operations
type databaseHandler interface {
	CreateFileMetaData(table string, fd *fileData)
}

// DB encapsulates a db connection with the operations
type DB struct {
	*sql.DB
}

// NewDB will create a new database connection with the supplied string
func NewDB(driver, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// CreateFileMetaData will take the data from the stored file
// and persist it to the database
func (db *DB) CreateFileMetaData(table string, fd *fileData) {
	sqlStatement := `INSERT INTO ` + table  + ` (name, slug, format, path, size)
	VALUES ($1, $2, $3, $4, $5)`
	if _, err := db.Exec(sqlStatement, fd.name, fd.slug, fd.format, fd.path, fd.size); err != nil {
		panic(err)
	}
}
