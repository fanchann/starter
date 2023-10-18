package main

import (
	"flag"
	"fmt"
	"starter/internal/config"
)

var conf config.IConfig
var configurationFile *string

func init() {
	configurationFile = flag.String("c", ".env", "configuration file not found!")
	flag.Parse()
	conf = config.NewLoadConfig(*configurationFile)
}

func main() {
	// cfg := config.NewLoadEnvConfig()
	// db := config.NewPostgresDatabaseConnection().Connect(cfg)
	// db.AutoMigrate(&users{})
	// d, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }

	// err2 := d.Ping()
	// if err2 != nil {
	// 	panic(err2)
	// }

	// fmt.Println("success conect")

	// cfg := config.NewLoadEnvConfig()
	// fmt.Printf("cfg.Get(\"DB_USERNAME\"): %v\n", cfg.Get("DB_USERNAME"))

	//conf := config.NewLoadConfig(".env")
	fmt.Printf("conf.Get(\"APP_NAME\"): %v\n", conf.Get("APP_NAME"))

}
