package main

import (
	"context"
	"fmt"
	"ginMongo/controllers"
	"ginMongo/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	userservice services.UserService
	usercontroller controllers.Usercontroller
	ctx context.Context
	usercollection *mongo.Collection
	mongoclient *mongo.Client
	err error
)

func init()  {
	//create context with no cancelation
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}

	// use admin command to check mongodb
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection has established")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection , ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main()  {
	// disconnect from mongo if connection shutdown
	defer mongoclient.Disconnect(ctx)

	// v1/user/create
	basePath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}