package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	// It's used on sql.Open().
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersPort     = "mysql_users_port"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = goDotEnvVariable(mysqlUsersUsername)
	password = goDotEnvVariable(mysqlUsersPassword)
	host     = goDotEnvVariable(mysqlUsersHost)
	port     = goDotEnvVariable(mysqlUsersPort)
	schema   = goDotEnvVariable(mysqlUsersSchema)
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Printf("Database successfully connected to %s", host)
}
