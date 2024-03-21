package controller

import (
	"blog/db"
	"blog/helper"
	"blog/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBlogs(c *gin.Context) {
	_, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	blogs := []model.Blog{}
	db.Database.Find(&blogs)

	c.JSON(200, gin.H{"blogs": blogs})
}

func GetBlogsByUser(c *gin.Context){
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	blogs:= user.Blogs
	c.JSON(200, gin.H{"blogs": blogs})
}

func GetBlogById(c *gin.Context) {
	blog := model.Blog{}
	db.Database.Preload("User").Preload("Comments").First(&blog, c.Param("id"))
	c.JSON(200, gin.H{"blog": blog})
}

func CreateBlog(c *gin.Context){
	var input model.Blog

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID
	savedTodo, err := input.Save()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"blog": savedTodo})
}

func UpdateBlog(c *gin.Context){
	id:=c.Param("id")

	blogID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	existingBlog := model.Blog{}
	existingBlog.UserID = user.ID
	existingBlog.ID = uint(blogID)

	if result := db.Database.First(&existingBlog, blogID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&existingBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := existingBlog.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": existingBlog})
}

func DeleteBlog(context *gin.Context){
	id:=context.Param("id")

	todoId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingBlog := model.Blog{}
	existingBlog.UserID = user.ID
	existingBlog.ID = uint(todoId)

	if result:=db.Database.First(&existingBlog); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if err := existingBlog.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}