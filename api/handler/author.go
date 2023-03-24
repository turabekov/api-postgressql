package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAuthor(c *gin.Context) {

	var createAuthor models.CreateAuthor

	err := c.ShouldBindJSON(&createAuthor) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create author", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Author().Create(&createAuthor)
	if err != nil {
		h.handlerResponse(c, "storage.author.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Book().GetByID(&models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.author.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create author", http.StatusCreated, resp)
}

func (h *Handler) GetByIdAuthor(c *gin.Context) {
	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id author", http.StatusBadRequest, "invalid author id")
		return
	}

	resp, err := h.storages.Author().GetByID(&models.AuthorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.author.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get author by id", http.StatusCreated, resp)
}

func (h *Handler) GetListAuthor(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list author", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list author", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Author().GetList(&models.GetListAuthorRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.author.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list author response", http.StatusOK, resp)
}

func (h *Handler) UpdateAuthor(c *gin.Context) {

	var updateAuthor models.UpdateAuthor

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id author", http.StatusBadRequest, "invalid author id")
		return
	}

	err := c.ShouldBindJSON(&updateAuthor)
	if err != nil {
		h.handlerResponse(c, "update author", http.StatusBadRequest, err.Error())
		return
	}

	updateAuthor.Id = id

	rowsAffected, err := h.storages.Author().Update(&updateAuthor)
	if err != nil {
		h.handlerResponse(c, "storage.author.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.author.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Author().GetByID(&models.AuthorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.author.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update author", http.StatusAccepted, resp)
}

func (h *Handler) DeleteAuthor(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id author", http.StatusBadRequest, "invalid author id")
		return
	}

	rowsAffected, err := h.storages.Author().Delete(&models.AuthorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.author.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.author.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "update author", http.StatusAccepted, nil)
}
