package connection

import (
	"apigolang/src/apigo/utils"
	DBConfig "apigolang/src/config/db"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type Connection interface {
	SqlDb() *sql.DB
}

type Postgres struct {
	db *sql.DB
}

var (
	dbPG   *Postgres
	connPG *sql.DB
)

func InitCloseConnection() {
	if err := closePostgresConnection(); err != nil {
		utils.LogError(err, "failed to close PostgresSQL connection")
	}
}

func (p *Postgres) SqlDb() *sql.DB {
	return p.db
}

func NewPostgresConnection(config *DBConfig.Config) (*Postgres, error) {
	conn, err := sql.Open("postgres", config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening PostgreSQL connection: %v", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging PostgreSQL connection: %v", err)
	}

	if config.UseConnectionPool {
		conn.SetMaxOpenConns(config.MaxOpenCons)
		conn.SetMaxIdleConns(config.MaxIdleCons)
		conn.SetConnMaxIdleTime(config.ConnMaxIdleTime)
		connPG = conn
	}

	dbPG = &Postgres{db: conn}
	return dbPG, nil
}

func closePostgresConnection() error {
	if connPG != nil {
		if err := connPG.Close(); err != nil {
			return fmt.Errorf("failed to close PostgreSQL connection: %v", err)
		}
		utils.LogSuccess("Closing Postgres connection...", "Closing Postgres connection...")
	}
	return nil
}
