package postgre

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewDBConnection(driver, host string) (conn *sql.DB) {
	conn, err := sql.Open(driver, host)
	if err != nil {
		log.Println("Failed db : ", err)
	} else {
		conn.SetConnMaxLifetime(2 * time.Second)
	}
	return
}
