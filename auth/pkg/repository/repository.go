package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// Repository - repo
type Repository struct {
	Authorization
}

// NewRepository - constructor
func NewRepository(db *sqlx.DB) *Repository {
	createAccountTable(db)
	return &Repository{
		Authorization: NewAuth(db),
	}
}

func createAccountTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"token" TEXT,
		"username" TEXT,		
		"password_hash" TEXT		
	  );`

	fmt.Println("Create auth.account table...")
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create auth.account table error", err.Error())
	}
	statement.Exec()
	fmt.Println("auth.account table created")
}
