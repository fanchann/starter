package main

import (
	"fmt"
	"starter/internal/config"
)

type users struct {
	Name  string
	Email string
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

	cfg := config.NewLoadloadConfig("config.toml")
	fmt.Printf("cfg.Get(\"APP_NAME\"): %v\n", cfg.Get("APP_NAME"))

}
