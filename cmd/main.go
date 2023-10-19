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
	db := config.NewMysqlDatabaseConnection().Connect(conf)
	d, err := db.DB()
	if err != nil {
		panic(err)
	}

	err2 := d.Ping()
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("success connect!")

}
