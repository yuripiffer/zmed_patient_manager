package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
	"zmed_patient_manager/infrastructure/config"
)

func InitDBConnection() *pgxpool.Pool {
	connConfig, err := pgxpool.ParseConfig(getDBConnectionString())
	if err != nil {
		panic(err)
	}

	connConfig.MinConns = 5
	connConfig.MaxConns = 20
	connConfig.MaxConnLifetime = time.Hour * 1
	connConfig.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		panic(err)
	}
	return pool
}

func CloseDBConnection(dbConnection *pgxpool.Pool) {
	if dbConnection != nil {
		dbConnection.Close()
	}
}

func getDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.ENV.DBHost,
		config.ENV.DBPort,
		config.ENV.DBUser,
		config.ENV.DBName,
		config.ENV.DBPassword)
}
