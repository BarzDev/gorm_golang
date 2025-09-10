package controller

import (
	"net/http"
	"strconv"

	"library-api/model"
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
	a.rg.POST("/authors", a.create)
	a.rg.PUT("/authors/:id", a.update)
	a.rg.DELETE("/authors/:id", a.delete)
}

func (a *AuthorController) listAuthors(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "7"))

	authors, paging, err := a.authorUC.GetAll(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get authors")
		return
	}

	common.SendPagedResponse(c, authors, paging, "success")
}

func (a *AuthorController) getById(c *gin.Context) {
	id := c.Param("id")

	author, err := a.authorUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to get author id "+id+"")
		return
	}

	common.SendSingleResponse(c, author, "success")
}

func (a *AuthorController) create(c *gin.Context) {
	var payload model.AuthorRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	author, err := a.authorUC.Create(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to create author")
		return
	}

	common.SendSingleResponse(c, author, "success")
}

func (a *AuthorController) update(c *gin.Context) {
	id := c.Param("id")
	var payload model.AuthorRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	author, err := a.authorUC.Update(id, payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to update author id "+id+"")
		return
	}

	common.SendSingleResponse(c, author, "success")
}

func (a *AuthorController) delete(c *gin.Context) {
	id := c.Param("id")

	if err := a.authorUC.Delete(id); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to delete author id "+id+"")
		return
	}

	common.SendSingleResponse(c, nil, "success")
}
