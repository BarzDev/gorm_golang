package controller

import (
	"net/http"
	"strconv"

	"library-api/shared/common"
	"library-api/usecase"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUseCase
	rg       *gin.RouterGroup
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.RouterGroup) *AuthorController {
	return &AuthorController{
		authorUC: authorUC,
		rg:       rg,
	}
}

func (a *AuthorController) Route() {
	a.rg.GET("/authors", a.listAuthors)
	a.rg.GET("/authors/:id", a.getById)
}

func (a *AuthorController) listAuthors(c *gin.Context) {
	authors, err := a.authorUC.GetAll()
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get authors")
		return
	}

	common.SendSingleResponse(c, authors, "success")
}

func (a *AuthorController) getById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid id format")
		return
	}

	author, err := a.authorUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get author")
		return
	}

	common.SendSingleResponse(c, author, "success")
}
