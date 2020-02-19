package sql

import (
	"github.com/Chewy-Inc/lets-go/v1/util"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	connString = os.Getenv("POSTGRES_URL")
	conn, _    = OpenConnection(connString)
	log, _ = util.InitLoggerWithLevel(nil)
)

//Opens a postgres connection using the SQL ORM library Gorm. Example connString
//format:
//postgresql://localhost:5432/denver_replica?sslmode=disable
func OpenConnection(connString string) (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panic(err)
	}
	conn.SingularTable(true)
	return conn, err
}