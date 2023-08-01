package controllers

import (
	"ginMongo/models"
	"ginMongo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	UserService services.UserService
}

// The constractor
func New(userservice services.UserService) Usercontroller {
	return Usercontroller{
		UserService: userservice,
	}
}

func (uc *Usercontroller) CreateUser(ctx *gin.Context) {

	// ctx hole the user information
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *Usercontroller) GetUser(ctx *gin.Context)  {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *Usercontroller) GetAll(ctx *gin.Context)  {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *Usercontroller) UpdateUser(ctx *gin.Context) {
	// ctx hole the user information
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *Usercontroller) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Define the routes
func (uc *Usercontroller) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create" , uc.CreateUser)
	userroute.GET("/get/:name" , uc.GetUser)
	userroute.GET("/getall" , uc.GetAll)
	userroute.PATCH("/update" , uc.UpdateUser)
	userroute.DELETE("/delete/:name" , uc.DeleteUser)
}