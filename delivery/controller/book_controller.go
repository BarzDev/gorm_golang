package controller

import (
	"net/http"
	"strconv"

	"library-api/model"
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
	b.rg.GET("/books/list", b.listBooks)
	b.rg.GET("/books/:id", b.getById)
}

func (b *BookController) listBooks(c *gin.Context) {
	var authorID, categoryID *int

	if a := c.Query("author_id"); a != "" {
		id, err := strconv.Atoi(a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "invalid author_id",
			})
			return
		}
		authorID = &id
	}

	if k := c.Query("category_id"); k != "" {
		id, err := strconv.Atoi(k)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "invalid category_id",
			})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "failed to get books",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    books,
	})
}

func (b *BookController) getById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	book, err := b.bookUC.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "failed to get book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    book,
	})
}
