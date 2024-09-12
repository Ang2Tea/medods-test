package common

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (c *DBConfig) DSN() string {
	params := map[string]string{
		"host":     c.Host,
		"port":     c.Port,
		"user":     c.Username,
		"password": c.Password,
		"dbname":   c.Database,
		"sslmode":  "disable",
	}

	stringBuild := make([]string, 0, len(params))
	for key, value := range params {
		stringBuild = append(stringBuild, key+"="+value)
	}

	return strings.Join(stringBuild, " ")
}

func PostgresGormConnect(host, port, username, password, database string) (*gorm.DB, error) {
	cfg := DBConfig{
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
	}

	if err := createIfNotExistDatabase(cfg); err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_, err = db.DB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createIfNotExistDatabase(cfg DBConfig) error {
	const (
		sqlQueryExistDatabase = `SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)`
		// В Exec нужно добавить database name
		sqlQueryCreateDatabase = `CREATE DATABASE `
	)

	const defaultPostgresDatabase = "postgres"
	cfgDBPostgres := DBConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: defaultPostgresDatabase,
	}

	db, err := gorm.Open(postgres.Open(cfgDBPostgres.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	var dbExist bool

	tx := db.Raw(sqlQueryExistDatabase, cfg.Database).Scan(&dbExist)
	if tx.Error != nil {
		return tx.Error
	}

	if !dbExist {
		tx = db.Exec(sqlQueryCreateDatabase + cfg.Database)
		if tx.Error != nil {
			return tx.Error
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	var host, username, password, database, port string

	LookupEnv(&host, POSTGRES_HOST, "localhost")
	LookupEnv(&username, POSTGRES_USERNAME)
	LookupEnv(&password, POSTGRES_PASSWORD)
	LookupEnv(&database, POSTGRES_DATABASE_NAME)
	LookupEnv(&port, POSTGRES_PORT, "5432")

	db, err := PostgresGormConnect(host, port, username, password, database)
	if err != nil {
		panic(err.Error())
	}

	return db
}
