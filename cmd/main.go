package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"app/api"
	"app/config"
	"app/storage/postgresql"
)

func main() {
	cfg := config.Load()

	store, err := postgresql.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Println("Error connect to postgresql: ", err.Error())
		return
	}

	defer store.CloseDB()

	r := gin.New()

	api.NewApi(r, &cfg, store)

	fmt.Println("Server Listening port", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Println("Error listening server:", err.Error())
		return
	}
}
