package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabaseConnection interface {
	Connect(cfg IConfig) interface{}
}

type mongoDatabaseConnection struct{}

func (m mongoDatabaseConnection) Connect(cfg IConfig) interface{} {
	mongoUri := fmt.Sprintf("mongodb://%s:%s", cfg.Get("MONGO_HOST"), cfg.Get("MONGO_PORT"))
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctxWithTimeout)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewMongoDatabaseConnection() IDatabaseConnection {
	return &mongoDatabaseConnection{}
}

type mysqlDatabaseConnection struct{}

func (m *mysqlDatabaseConnection) Connect(cfg IConfig) interface{} {
	// get environment
	dbUsername := cfg.Get("DB_USERNAME")
	dbPassword := cfg.Get("DB_PASSWORD")
	dbHost := cfg.Get("DB_HOST")
	dbPort := cfg.Get("DB_PORT")
	dbName := cfg.Get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewMysqlDatabaseConnection() IDatabaseConnection {
	return &mysqlDatabaseConnection{}
}

type postgresDatabaseConnection struct{}

func (p *postgresDatabaseConnection) Connect(cfg IConfig) interface{} {
	// get environment
	dbUsername := cfg.Get("DB_USERNAME")
	dbPassword := cfg.Get("DB_PASSWORD")
	dbHost := cfg.Get("DB_HOST")
	dbPort := cfg.Get("DB_PORT")
	dbName := cfg.Get("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewPostgresDatabaseConnection() IDatabaseConnection {
	return &postgresDatabaseConnection{}
}
