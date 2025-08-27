package controller

import (
	"library-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUseCase
	rg       *gin.RouterGroup
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.RouterGroup,) *AuthorController{
	return &AuthorController{
		authorUC : authorUC,
		rg: rg,
		
	}
}

func (a *AuthorController) Route() {
	a.rg.GET("/authors/list", a.listAuthors)
	a.rg.GET("/authors/:id", a.getById)
}


func (a *AuthorController) listAuthors(c *gin.Context) {
	authors, err := a.authorUC.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": "failed to get authors",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message":"success",
		"data": authors,
	})
}

func (a *AuthorController) getById(c *gin.Context){

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}
	
	author, err := a.authorUC.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"error": "failed to get author",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message":"success",
		"data": author,
	})
}