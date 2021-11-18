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
		"public_id" TEXT NOT NULL,
		"name" TEXT NOT NULL,
		"token" TEXT,
		"username" TEXT NOT NULL,		
		"password_hash" TEXT NOT NULL,
		"role" INTEGER DEFAULT 0,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
	  );`

	fmt.Println("Create billing.account table...")
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create billing.account table error", err.Error())
	}
	statement.Exec()
	fmt.Println("billing.account table created")
}
