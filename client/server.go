package client

import (
	"github.com/gin-gonic/gin"

	"github.com/go-bp-ikm-api/config"
	"github.com/go-bp-ikm-api/docs"

	"github.com/go-bp-ikm-api/middlewares"
	"github.com/go-bp-ikm-api/routes"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	rc routes.CategoryRoute
)

type Server struct {
	config   config.Config
	router   *gin.Engine
	database *mongo.Database
}

func NewServer(config config.Config, mongoconn *mongo.Client) (*Server, error) {

	server := &Server{
		config:   config,
		router:   gin.Default(),
		database: mongoconn.Database(config.DATABASE),
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	categoryc := server.database.Collection("category")
	rc = routes.NewCategoryRoute(categoryc)

	server.router = gin.Default()
	server.router.Use(middlewares.CORS())

}

func (server *Server) Start(address string) error {
	basepath := server.router.Group(docs.SwaggerInfo.BasePath)
	{
		rc.RegisterCategoryRoutes(basepath)
	}

	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// rc.RegisterCategoryRoutes(basepath)
	return server.router.Run(address)
}
