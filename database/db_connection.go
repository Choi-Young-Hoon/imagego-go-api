package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"imagego-go-api/util"
)

var defaultDBConnection *DBConnection

func GeteDefaultDBConnection() *DBConnection {
	return defaultDBConnection
}

func CreateDefaultDBConnection(databaseConfig util.DatabaseConfig) *DBConnection {
	connection := &DBConnection{
		Host:     databaseConfig.Host,
		User:     databaseConfig.User,
		Password: databaseConfig.Password,
		DBName:   databaseConfig.DBName,
		Port:     databaseConfig.Port,
	}

	defaultDBConnection = connection

	return connection
}

type DBConnection struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	db       *gorm.DB
}

func (dbs *DBConnection) Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbs.Host, dbs.User, dbs.Password, "postgres", dbs.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbs.db = db

	return nil
}

func (dbs *DBConnection) Close() {
	db, err := dbs.db.DB()
	if err != nil {
		panic(err)
	}

	db.Close()
}

func (dbs *DBConnection) Create(data interface{}) error {
	result := dbs.db.Create(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dbs *DBConnection) Update(data interface{}) error {
	result := dbs.db.Save(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dbs *DBConnection) Delete(data interface{}) error {
	result := dbs.db.Delete(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dbs *DBConnection) Select(data interface{}) error {
	// gorm 에서 Find 시에 data 의 값을 기준으로 검색하는방법 알려줘

	result := dbs.db.Find(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dbs *DBConnection) GetDB() *gorm.DB {
	return dbs.db
}
