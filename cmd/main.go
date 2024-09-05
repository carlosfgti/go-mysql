package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "username"),
		Password: getEnv("DB_PASSWORD", "userpass"),
		Database: getEnv("DB_DATABASE", "goproject"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", LoadDatabaseConfig().User, LoadDatabaseConfig().Password, LoadDatabaseConfig().Host, LoadDatabaseConfig().Port, LoadDatabaseConfig().Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully connected! \n")

	selectRows(db)
}

func selectRows(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
}
