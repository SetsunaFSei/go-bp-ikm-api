package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-bp-ikm-api/client"
	"github.com/go-bp-ikm-api/config"
	"github.com/go-bp-ikm-api/docs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server      *gin.Engine
	ctx         *gin.Context
	userc       *mongo.Collection
	categoryc   *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mongoconn := options.Client().ApplyURI(config.DBSource)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	runGinServer(config, mongoclient)
}

func runGinServer(config config.Config, mongoconn *mongo.Client) {
	server, err := client.NewServer(config, mongoclient)
	docs.SwaggerInfo.Title = "Celestial - Celestial API"
	docs.SwaggerInfo.Description = "Celestial - Celestial App API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.HTTPServerAddress
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
