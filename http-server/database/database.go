package database

import (
	"database/sql"
	"fmt"

	"github.com/cenkalti/backoff"
	"github.com/swisslockTech/go-http/http-server/logging"
)

const (
	// no tracing
	dbConnectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbDriverName             = "postgres"

	defaultPingMaxRetry = 10
)

func New() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector")

	cfg := loadConfig()

	connString := fmt.Sprintf(dbConnectionStringFormat,
		cfg.dbHost, cfg.dbPort,
		cfg.dbUsername, cfg.dbPassword, cfg.dbName,
		cfg.dbSslMode,
	)
	logging.SugaredLog.Debugf("DB connection string: %s", connString)

	return sql.Open(dbDriverName, connString)
}

func InitDb(db *sql.DB) error {
	logging.Log.Info("Initialize DB")

	result, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return tableErr
	}
	logging.SugaredLog.Debugf("Initializion result: %s", result)
	return nil
}

func PingDb(db *sql.DB, maxRetry uint64) error {
	if maxRetry <= 0 {
		logging.SugaredLog.Warnf("PingDB maxRetry value not valid, falling back to default (%d)", defaultPingMaxRetry)
		maxRetry = defaultPingMaxRetry
	}

	// WARN: connection takes a bit time to be opened, golang application is so fast that the first ping could easily fail
	pingErr := backoff.Retry(
		func() error {
			err := db.Ping()
			if err != nil {
				logging.Log.Info("PostgreSQL connection not ready, backing off...")
				return err
			}
			logging.Log.Info("PostgreSQL connection ready")
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry),
	)
	return pingErr
}
