package main

import (
	"fmt"
	"log"
	"net/http"

	"app/config"
	"app/controller"
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

	newController := controller.NewController(&cfg, store)

	http.HandleFunc("/book", newController.BookController)
	http.HandleFunc("/author", newController.AuthorController)


	fmt.Println("Listening Server", cfg.ServerHost+cfg.ServerPort)
	err = http.ListenAndServe(cfg.ServerHost+cfg.ServerPort, nil)
	if err != nil {
		log.Println("Error listening server:", err.Error())
		return
	}
}
