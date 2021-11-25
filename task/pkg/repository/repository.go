package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// Repository - repo
type Repository struct {
	Authorizer
	Tasker
}

// NewRepository - constructor
func NewRepository(db *sqlx.DB) *Repository {
	createAccountTable(db)
	createTaskTable(db)

	return &Repository{
		Authorizer: NewAuth(db),
		Tasker:     NewTask(db),
	}
}

func createAccountTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,			
		"public_id" TEXT NOT NULL,
		"name" TEXT NOT NULL,
		"token" TEXT,
		"username" TEXT NOT NULL,
		"role" INTEGER DEFAULT 0,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
	  );`

	fmt.Println("Create task.account table...")
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create task.account table error", err.Error())
	}
	statement.Exec()
	fmt.Println("task.account table created")
}

func createTaskTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS task (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"assigned_account_id" INTEGER NOT NULL,		
		"public_id" TEXT NOT NULL,
		"description" TEXT,
		"jira_id" TEXT,
		"status" INTEGER DEFAULT 0,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		FOREIGN KEY(assigned_account_id) REFERENCES account(public_id)
	  );`

	fmt.Println("Create task.task table...")
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create task.task table error", err.Error())
	}
	statement.Exec()
	fmt.Println("task.task table created")
}
