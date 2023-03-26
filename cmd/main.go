package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"app/api"
	"app/config"
	"app/pkg/logger"
	"app/storage/postgresql"
)

func main() {
	cfg := config.Load()

	// loger start-----------------------------------------------------------------------
	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("app", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()
	// loger end-----------------------------------------------------------------------

	store, err := postgresql.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Panic("Error connect to postgresql: ", logger.Error(err))
		return
	}

	defer store.CloseDB()

	r := gin.New()
	// call logger
	r.Use(gin.Recovery(), gin.Logger())
	// cal api
	api.NewApi(r, &cfg, store)

	// running server
	fmt.Println("Server Listening port", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server:", logger.Error(err))
		return
	}
}
