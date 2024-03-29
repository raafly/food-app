// MAKE SURE YOU CHANGE THE URL POSTGRES

package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/raafly/food-app/pkg/helpers"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:saturna@localhost:5432/food?sslmode=disable")
	helper.PanicIfError(err)
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100) 
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}