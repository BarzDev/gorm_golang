package controller

import (
	"library-api/usecase"
	"net/http"

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