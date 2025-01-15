package connectors
import (
"database/sql"
"fmt"
  _ "github.com/go-sql-driver/mysql"
)
type Database struct {
DB *sql.DB
}
func NewDatabase() *Database {
// Create a new database configuration
cfg := newDatabaseCfg()
// Create a connection string
connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
// Open a new connection to the database
db, err := sql.Open("mysql", connectionString)
if err != nil {
	panic(err)
	}
	return &Database{
	DB: db,
	}
	}