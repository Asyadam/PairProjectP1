package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func InitDB() (db *sql.DB, err error) {

	db, err = sql.Open(
		"mysql",
		os.Getenv("MYSQL_DSN"),
	)

	if err != nil {
		fmt.Println(err)
		return db, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return db, err
	}

	fmt.Println("Database connected!")

	return db, nil
}
