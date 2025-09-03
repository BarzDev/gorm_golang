package controller

import (
	"net/http"

	"library-api/model"
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
	c.rg.POST("/categories", c.create)
	c.rg.PUT("/categories/:id", c.update)
	c.rg.DELETE("/categories/:id", c.delete)
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
	id := ctx.Param("id")
	category, err := c.categoryUC.GetById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to get category")
		return
	}

	common.SendSingleResponse(ctx, category, "success")
}

func (c *CategoryController) create(ctx *gin.Context) {
	var payload model.CategoryRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	category, err := c.categoryUC.Create(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to create category")
		return
	}

	common.SendSingleResponse(ctx, category, "success")
}

func (c *CategoryController) update(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload model.CategoryRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	category, err := c.categoryUC.Update(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to update category id "+id+"")
		return
	}

	common.SendSingleResponse(ctx, category, "success")
}

func (c *CategoryController) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.categoryUC.Delete(id); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "failed to delete category id "+id+"")
		return
	}

	common.SendSingleResponse(ctx, nil, "success")
}
