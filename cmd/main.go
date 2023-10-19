package main

import (
	"context"
	"flag"
	"log"
	"starter/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var conf config.IConfig
var configurationFile *string

func init() {
	configurationFile = flag.String("c", ".env", "configuration file not found!")
	flag.Parse()
	conf = config.NewLoadConfig(*configurationFile)
}

func main() {
	// db := config.NewMysqlDatabaseConnection().Connect(conf).(*gorm.DB)
	// d, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }

	// err2 := d.Ping()
	// if err2 != nil {
	// 	panic(err2)
	// }
	//fmt.Println("success connect!")

	mongo := config.NewMongoDatabaseConnection().Connect(conf).(*mongo.Client)
	err := mongo.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

}
