package mysqlunclosed

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Run opens a MySQL connection and intentionally does not close rows/db.
func Run() {
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("[mysql_unclosed] open error: %v\n", err)
		return
	}

	// Bug pattern: db.Close() is missing, causing connection/resource leak.
	rows, err := db.Query("SELECT 1")
	if err != nil {
		fmt.Printf("[mysql_unclosed] query error: %v\n", err)
		return
	}

	// Bug pattern: rows.Close() is also missing.
	_ = rows
	fmt.Println("[mysql_unclosed] query executed without closing resources")
}
