package api

import (
	"app/api/handler"
	"app/config"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI) {
	handler := handler.NewHandler(cfg, store)

	// book api
	r.POST("/book", handler.CreateBook)
	r.GET("/book/:id", handler.GetByIdBook)
	r.GET("/book", handler.GetListBook)
	r.PATCH("/book/:id", handler.UpdateBook)
	r.DELETE("/book/:id", handler.DeleteBook)

	// author api
	r.POST("/author", handler.CreateAuthor)
	r.GET("/author/:id", handler.GetByIdAuthor)
	r.GET("/author", handler.GetListAuthor)
	r.PUT("/author/:id", handler.UpdateAuthor)
	r.DELETE("/author/:id", handler.DeleteAuthor)

	// user api
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

}
