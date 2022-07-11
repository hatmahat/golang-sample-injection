package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Config struct {
	Db *sqlx.DB
}

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type dbConfig struct {
	dbHost     string
	dbPort     string
	dbName     string
	dbUser     string
	dbPassword string
	dbDriver   string
}

func (c *Config) initDb() {
	var dbConfig = dbConfig{}
	dbConfig.dbHost = os.Getenv("DB_HOST")
	dbConfig.dbPort = os.Getenv("DB_PORT")
	dbConfig.dbUser = os.Getenv("DB_USER")
	dbConfig.dbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.dbName = os.Getenv("DB_NAME")
	//dbConfig.dbDriver = os.Getenv("DB_DRIVER")
	dbConfig.dbDriver = "postgres"

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.dbDriver,
		dbConfig.dbUser,
		dbConfig.dbPassword,
		dbConfig.dbHost,
		dbConfig.dbPort,
		dbConfig.dbName,
	)

	db, err := sqlx.Connect(dbConfig.dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	c.Db = db
}

func (c *Config) DbConn() *sqlx.DB {
	return c.Db
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}

/*
set DB_DRIVER=postgres
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASSWORD=87654321
set DB_NAME=db_credential
set API_HOST=localhost
set API_PORT=8888
*/
