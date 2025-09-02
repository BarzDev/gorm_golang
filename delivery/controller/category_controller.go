package controller

import (
	"net/http"
	"strconv"

	"library-api/shared/common"
	"library-api/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryUC usecase.CategoryUseCase
	rg         *gin.RouterGroup
}

func NewCategoryController(categoryUC usecase.CategoryUseCase, rg *gin.RouterGroup) *CategoryController {
	return &CategoryController{
		categoryUC: categoryUC,
		rg:         rg,
	}
}

func (c *CategoryController) Route() {
	c.rg.GET("/categories", c.listCategories)
	c.rg.GET("/categories/:id", c.getById)
}

func (c *CategoryController) listCategories(ctx *gin.Context) {
	categories, err := c.categoryUC.GetAll()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to get categories")
		return
	}

	common.SendSingleResponse(ctx, categories, "success")
}

func (c *CategoryController) getById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid id format")
	}

	category, err := c.categoryUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to get category")
		return
	}

	common.SendSingleResponse(ctx, category, "success")
}
