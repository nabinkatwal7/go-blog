package controller

import (
	"blog/helper"
	"blog/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input model.AuthenticationInputRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		ConfirmPassword: input.ConfirmPassword,
	}

	savedUser, err := user.Save()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, savedUser)
}

func Login(c *gin.Context) {
	var input model.AuthenticationInputLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByEmail(input.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(*user)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt": jwt,
	})
	
}