package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorizer
	Tasker
	Biller
}

func NewRepository(db *sqlx.DB) *Repository {
	createAccountTable(db)
	createTaskTable(db)
	createBillTable(db)
	createPaymentTable(db)

	return &Repository{
		Authorizer: NewAuth(db),
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

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create billing.account table error", err.Error())
	}
	statement.Exec()
	fmt.Println("billing.account table created 🗂")
}

func createTaskTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS task (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"assigned_account_id" TEXT NOT NULL,		
		"public_id" TEXT NOT NULL,
		"description" TEXT,
		"jira_id" TEXT,
		"status" INTEGER DEFAULT 0,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		FOREIGN KEY(assigned_account_id) REFERENCES account(public_id)
	  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create billing.task table error", err.Error())
	}
	statement.Exec()
	fmt.Println("billing.task table created 🗂")
}

func createBillTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS bill (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"public_id" TEXT NOT NULL,
		"account_id" TEXT		
		"task_id" TEXT,
		"transaction_reason" INTEGER,
		"price" INTEGER,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		FOREIGN KEY(account_id) REFERENCES account(public_id)
	  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create billing.bill table error", err.Error())
	}
	statement.Exec()
	fmt.Println("billing.bill table created 🗂")
}

func createPaymentTable(db *sqlx.DB) {
	query := `CREATE TABLE IF NOT EXISTS payment (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"public_id" TEXT NOT NULL,
		"account_id" TEXT		
		"amount" INTEGER,
		"created_at" DATE DEFAULT CURRENT_DATE NOT NULL,
		FOREIGN KEY(account_id) REFERENCES account(public_id)
	  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal("create billing.payment table error", err.Error())
	}
	statement.Exec()
	fmt.Println("billing.payment table created 🗂")
}
