package controllers

import (
	"net/http"

	"crud.com/api/models"
	"crud.com/api/services"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	UserService services.EmpService
}

func New(userservice services.EmpService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var emp models.Employee
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&emp)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Created"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	var empname string = ctx.Param("emp_name")
	user, err := uc.UserService.GetUser(&empname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	emps, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, emps)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var emp models.Employee
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&emp)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Updated"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var empname string = ctx.Param("emp_name")
	err := uc.UserService.DeleteUser(&empname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Sccessfully Deleted"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/employee")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:emp_name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:emp_name", uc.DeleteUser)
}