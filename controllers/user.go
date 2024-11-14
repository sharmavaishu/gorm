package controllers

import (
	"gorm_crud/database"
	"gorm_crud/models"
	"net/http"
	"strconv"
    "gorm.io/gorm"
	"github.com/gin-gonic/gin"
) 

type userRepo struct {
	Db *gorm.DB
}

func New() *userRepo {
	db := database.InitializeDatabase()

	// AutoMigrate will create the necessary tables based on the
   // defined models/structs
	db.AutoMigrate(&models.User{})
	return &userRepo{Db:db}
}

// get users

func (repository *userRepo) GetUsers(c *gin.Context){
        var user models.User
		err := models.GetUsers(repository.Db,&user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, user)
}


// get user by id

func (repository *userRepo) GetUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUser(repository.Db,&user,id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, user)
}

// create user

func (repository *userRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}


// update user 

func (repository *userRepo) UpdateUser(c *gin.Context) {
	var user models.User

	//get by id
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	//create
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}


// delete
func (repository *userRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}