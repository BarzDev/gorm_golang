package controller

import (
	"net/http"
	"strconv"

	"library-api/model"
	"library-api/shared/common"
	"library-api/shared/shared_model"
	"library-api/usecase"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUC usecase.BookUseCase
	rg     *gin.RouterGroup
}

func NewBookConroller(bookUC usecase.BookUseCase, rg *gin.RouterGroup) *BookController {
	return &BookController{
		bookUC: bookUC,
		rg:     rg,
	}
}

func (b *BookController) Route() {
	b.rg.GET("/books", b.listBooks)
	b.rg.GET("/books/:id", b.getById)
	b.rg.POST("/books", b.create)
	b.rg.PUT("/books/:id", b.update)
	b.rg.DELETE("/books/:id", b.delete)
}

func (b *BookController) listBooks(c *gin.Context) {
	var authorID, categoryID *int

	if a := c.Query("author_id"); a != "" {
		id, err := strconv.Atoi(a)
		if err != nil {
			common.SendErrorResponse(c, http.StatusBadRequest, "invalid author_id")
			return
		}
		authorID = &id
	}

	if k := c.Query("category_id"); k != "" {
		id, err := strconv.Atoi(k)
		if err != nil {
			common.SendErrorResponse(c, http.StatusBadRequest, "invalid category_id")
			return
		}
		categoryID = &id
	}

	var (
		books  []model.Book
		paging shared_model.Paging
		err    error
	)

	if authorID != nil || categoryID != nil {
		books, err = b.bookUC.Filter(authorID, categoryID)
	} else {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(c.DefaultQuery("size", "7"))
		books, paging, err = b.bookUC.GetAll(page, size)
	}

	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get books")
		return
	}

	common.SendPagedResponse(c, books, paging, "success")
}

func (b *BookController) getById(c *gin.Context) {
	id := c.Param("id")

	book, err := b.bookUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get book")

		return
	}

	common.SendSingleResponse(c, book, "success")
}

func (b *BookController) create(c *gin.Context) {
	var payload model.BookRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := b.bookUC.Create(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, book, "success")
}

func (b *BookController) update(c *gin.Context) {
	id := c.Param("id")
	var payload model.BookRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updateBook, err := b.bookUC.Update(id, payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to update book id "+id+"")
		return
	}

	common.SendSingleResponse(c, updateBook, "success")
}

func (b *BookController) delete(c *gin.Context) {
	id := c.Param("id")

	if err := b.bookUC.Delete(id); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to delete book id "+id+"")
		return
	}

	common.SendSingleResponse(c, nil, "success")
}
