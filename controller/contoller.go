package controller

import (
	"app/config"
	"app/storage"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	store storage.StorageI
	cfg   *config.Config
}

type Data struct {
	Err string `json:"error"`
}

func NewController(cfg *config.Config, store storage.StorageI) *Controller {
	return &Controller{
		store: store,
		cfg:   cfg,
	}
}

func (c *Controller) HandleFuncResponse(w http.ResponseWriter, tag string, code int, message string) {

	var data Data = Data{
		Err: message,
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(tag, message)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(body)
}
