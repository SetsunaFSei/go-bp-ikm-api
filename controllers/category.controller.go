package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-bp-ikm-api/common"
	"github.com/go-bp-ikm-api/dto"
	"github.com/go-bp-ikm-api/services"
	"github.com/go-bp-ikm-api/validator"
)

type CategoryController struct {
	CategoryService services.ICategoryService
}

func NewCategoryController(categoryservice services.ICategoryService) CategoryController {
	return CategoryController{
		CategoryService: categoryservice,
	}
}

// Create Category
// @Summary      Add an CreateCategoryDto
// @Description Create Category
// @Tags         category
// @Param        account  body      dto.CreateCategoryDto  true  "Add category"
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /category/create [post]
func (uc *CategoryController) CreateCategory(ctx *gin.Context) {

	var category dto.CreateCategoryDto
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}

	err := validator.ValidateCreateCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = uc.CategoryService.CreateCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetAll categories
// @Summary List categories
// @Description Get all categories
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /category [get]
func (uc *CategoryController) GetAll(ctx *gin.Context) {
	categorys, err := uc.CategoryService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categorys)
}

// GetAll categories by id
// @Summary   category by id
// @Description Get   category by id
// @Param        id       path      string                  true  "Id"
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /category [get]
func (uc *CategoryController) GetCategoryById(ctx *gin.Context) {
	var id string = ctx.Param("id")
	category := uc.CategoryService.GetCategoryById(&id)

	ctx.JSON(http.StatusOK, category)
	// rsp := mappping.NewUserResponse(user)
}
