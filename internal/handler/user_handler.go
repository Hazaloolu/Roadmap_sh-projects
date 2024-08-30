package handler

import (
	"blog_api/internal/auth"
	"blog_api/internal/model"
	"blog_api/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if email already exists
	var existingUser model.User

	if err := storage.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// hash password

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword

	// save the user

	if err := storage.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created Succesfully !"})

}
