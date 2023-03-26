package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Book godoc
// @ID create_book
// @Router /book [POST]
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body models.CreateBook true "CreateBookRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateBook(c *gin.Context) {

	var createBook models.CreateBook

	err := c.ShouldBindJSON(&createBook) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create book", http.StatusBadRequest, err.Error())
		return
	}

	// check status and calc sellprice
	if createBook.ProfitStatus == "fixed" && createBook.ProfitPrice >= 0 {
		createBook.SellPrice = createBook.IncomePrice + createBook.ProfitPrice
	} else if createBook.ProfitStatus == "precent" && createBook.ProfitPrice >= 0 {
		createBook.SellPrice = createBook.IncomePrice + createBook.IncomePrice*createBook.ProfitPrice/100
	} else {
		h.handlerResponse(c, "create book", http.StatusBadRequest, "given status incorrect")
		return
	}

	id, err := h.storages.Book().Create(&createBook)
	if err != nil {
		h.handlerResponse(c, "storage.book.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Book().GetByID(&models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.book.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create book", http.StatusCreated, resp)
}

// Get By ID Book godoc
// @ID get_by_id_book
// @Router /book/{id} [GET]
// @Summary Get By ID Book
// @Description Get By ID Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdBook(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id book", http.StatusBadRequest, "invalid book id")
		return
	}

	resp, err := h.storages.Book().GetByID(&models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.book.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get book by id", http.StatusCreated, resp)
}

// Get List Book godoc
// @ID get_list_book
// @Router /book [GET]
// @Summary Get List Book
// @Description Get List Book
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListBook(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list book", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list book", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Book().GetList(&models.GetListBookRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.book.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list book response", http.StatusOK, resp)
}

// Update Book godoc
// @ID update_book
// @Router /book/{id} [PATCH]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.UpdateBook true "UpdateBookRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateBook(c *gin.Context) {

	// var updateBook models.UpdateBook
	var body map[string]interface{}

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id book", http.StatusBadRequest, "invalid book id")
		return
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handlerResponse(c, "update book", http.StatusBadRequest, err.Error())
		return
	}

	body["id"] = id

	rowsAffected, err := h.storages.Book().Update(&body)
	if err != nil {
		h.handlerResponse(c, "storage.book.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.book.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Book().GetByID(&models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.book.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update book", http.StatusAccepted, resp)
}

// DELETE Book godoc
// @ID delete_book
// @Router /book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.BookPrimaryKey true "DeleteBookRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteBook(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id book", http.StatusBadRequest, "invalid book id")
		return
	}

	rowsAffected, err := h.storages.Book().Delete(&models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.book.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.book.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "update book", http.StatusNoContent, nil)
}
