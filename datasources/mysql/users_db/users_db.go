package users_db

import (
	"database/sql"
	"fmt"
	"github.com/comfysweet/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?", username, password, host, schema)
	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	if err := mysql.SetLogger(logger.GetLogger()); err != nil {
		panic(err)
	}
	logger.Info("database successfully configured")
}
