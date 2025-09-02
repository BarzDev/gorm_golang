package controller

import (
	"net/http"
	"strconv"

	"library-api/model"
	"library-api/shared/common"
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
		books []model.Book
		err   error
	)

	if authorID != nil || categoryID != nil {
		books, err = b.bookUC.Filter(authorID, categoryID)
	} else {
		books, err = b.bookUC.GetAll()
	}

	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get books")
		return
	}

	common.SendSingleResponse(c, books, "success")
}

func (b *BookController) getById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid id format")
		return
	}

	book, err := b.bookUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get book")

		return
	}

	common.SendSingleResponse(c, book, "success")
}
