package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-bp-ikm-api/controllers"
	"github.com/go-bp-ikm-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	categoryc *mongo.Collection
	cc        controllers.CategoryController
)

type CategoryRoute struct {
	CategoryController controllers.CategoryController
}

func NewCategoryRoute(collection *mongo.Collection) CategoryRoute {

	// repository := repository.NewCategoryRepository(collection, ctx, log)
	var ctx *gin.Context
	var loger log.Logger
	service := services.NewCategoryService(ctx, collection, loger)
	cc = controllers.NewCategoryController(service)
	return CategoryRoute{
		CategoryController: cc,
	}
}

func (uc *CategoryRoute) RegisterCategoryRoutes(rg *gin.RouterGroup) {
	categoryroute := rg.Group("/category")
	{
		categoryroute.GET("/", uc.CategoryController.GetAll)
		categoryroute.POST("/create", uc.CategoryController.CreateCategory)
	}

	// categoryroute.GET("/:id", uc.CategoryController.GetCategoryById)

	// categoryroute.PATCH("/update", uc.CategoryController.UpdateCategory)
	// categoryroute.DELETE("/delete/:name", uc.CategoryController.DeleteCategory)
}
