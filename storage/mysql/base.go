package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"vps-provider/config"

	_ "github.com/go-sql-driver/mysql"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jmoiron/sqlx"
)

var log = logging.Logger("storage")

// DB reference to database
var DB *sqlx.DB

const (
	tableNameUser = "users"
)

const (
	maxOpenConnections = 60
	connMaxLifetime    = 120
	maxIdleConnections = 30
	connMaxIdleTime    = 20
)

func Init(cfg *config.Config) error {
	if cfg.DatabaseURL == "" {
		return fmt.Errorf("database url not setup")
	}

	db, err := sqlx.Connect("mysql", cfg.DatabaseURL)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	DB = db

	return initTables()
}

// initializes data tables.
func initTables() error {
	// init table
	tx, err := DB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			log.Errorf("InitTables Rollback err:%s", err.Error())
		}
	}()

	// Execute table creation statements
	tx.MustExec(fmt.Sprintf(cUsersTable, tableNameUser))

	return tx.Commit()
}
