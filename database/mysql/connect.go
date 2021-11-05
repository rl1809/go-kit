package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

// DBConn returns a MySQL connection pool.
func DBConnnect() (*sql.DB, error) {


	db, err := sql.Open("mysql", "nha:nha@/mydb")

	return db, err
}

