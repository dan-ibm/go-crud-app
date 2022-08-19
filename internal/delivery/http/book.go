package http

import (
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create Book
// @Security ApiKeyAuth
// @Tags books
// @Description create Book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body domain.BookInput true "book info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books [post]
func (h *Handler) createBook(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.BookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Book.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, statusResponse{"ok"})
}

type getAllBooksResponse struct {
	Data []domain.Book `json:"data"`
}

// @Summary Get AllBooks
// @Security ApiKeyAuth
// @Tags books
// @Description gets all books
// @ID get-all-books
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllBooksResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books [get]
func (h *Handler) getAllBooks(c *gin.Context) {

	books, err := h.services.Book.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

// @Summary Get BooksUser
// @Security ApiKeyAuth
// @Tags books
// @Description get books by user
// @ID get-book-by-user
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllBooksResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/user [get]
func (h *Handler) getAllBooksByUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	books, err := h.services.Book.GetAllByUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

// @Summary Get Book By Id
// @Security ApiKeyAuth
// @Tags books
// @Description get list by id
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} domain.Book
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/{id} [get]
func (h *Handler) getBookById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.services.Book.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary Update Book
// @Security ApiKeyAuth
// @Tags books
// @Description update Book
// @ID update-book
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/{id} [put]
func (h *Handler) updateBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.BookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Book.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Book
// @Security ApiKeyAuth
// @Tags books
// @Description delete Book
// @ID delete-book
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/books/{id} [delete]
func (h *Handler) deleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Book.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
