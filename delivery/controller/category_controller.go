package controller

import (
	"net/http"
	"strconv"

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
	c.rg.GET("/categories/list", c.listCategories)
	c.rg.GET("/categories/:id", c.getById)
}

func (c *CategoryController) listCategories(ctx *gin.Context) {
	categories, err := c.categoryUC.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "failed to get authors",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    categories,
	})
}

func (c *CategoryController) getById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	category, err := c.categoryUC.GetById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "failed to get category",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    category,
	})
}
