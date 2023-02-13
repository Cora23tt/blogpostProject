package main

import (
	"net/http"
	"strconv"
	"testrepo/database"

	_ "testrepo/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title BlogPost API Documentation
// @version 1.0
// @license.name Apache 2.0
// @description This is a sample demo-project
// @contact.email aziz.rustamov.mail@gmail.com
func main() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	database.ConnectDB()

	r.GET("/list", allPostsHandler)
	r.GET("/list/:num", listPostsHandler)

	r.GET("/post/:id", postHandler)
	r.POST("/post", createPostHandler)
	r.DELETE("/post/:id", deletePostHandler)
	r.POST("/post/:id", editPostHandler)

	r.POST("/bytag", getPostByTag)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":9000")
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

type ErrorMessage struct {
	Error string
}

// @Summary		Recive a list of all posts
// @Description	get all posts
// @Tags		Post
// @Accept		json
// @Produce		json
// @Success		200 {object} []database.Post
// @Failure		500 {object} ErrorMessage
// @Router		/list [get]
func allPostsHandler(c *gin.Context) {
	posts, err := database.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// @Summary		Recive a list of posts
// @Description	get list by limit {num}
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		num	path		int	true	"limit number of posts"
// @Success		200 {object} []database.Post
// @Failure		409 {object} ErrorMessage
// @Failure		400 {object} ErrorMessage
// @Router		/list/{num} [get]
func listPostsHandler(c *gin.Context) {
	s := c.Param("num")
	num, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var posts []database.Post
	posts, err = database.GetList(num)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// @Summary		Recive post by ID
// @Description	get post by ID {id}
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"id number of post"
// @Success		200 {object} database.Post
// @Failure		409 {object} ErrorMessage
// @Failure		400 {object} ErrorMessage
// @Router		/post/{id} [get]
func postHandler(c *gin.Context) {
	s := c.Param("id")
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post database.Post
	post, err = database.GetPost(n)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary		Creates a post
// @Description	Create new post
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		post	body		database.Post	true	"The post to be created"
// @Success		200 {number} int "the final ID of the post"
// @Failure		409 {object} ErrorMessage
// @Failure		400 {object} ErrorMessage
// @Router		/post [post]
func createPostHandler(c *gin.Context) {
	var newpost database.Post
	if err := c.BindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := database.AddPost(newpost)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

// @Summary		Deletes the post from DB
// @Description	delete post by ID
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"limit number of posts"
// @Success		200
// @Failure		409 {object} ErrorMessage
// @Failure		500 {object} ErrorMessage
// @Router		/post/{id} [delete]
func deletePostHandler(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = database.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// @Summary		Edits the post from DB by ID
// @Description	Edit post by ID
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		post	body		database.Post	true	"New post"
// @Param		id		path		int				true	"Specify the ID of the old post"
// @Success		200
// @Failure		400 {object} ErrorMessage
// @Failure		500 {object} ErrorMessage
// @Router		/post/{id} [post]
func editPostHandler(c *gin.Context) {
	var newpost database.Post
	if err := c.BindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.EditPost(newpost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Recives the list of posts by tag
// @Description	Get list by tag
// @Tags		Post
// @Accept		json
// @Produce		json
// @Param		tag	body		database.TagObj	true	"The tag of post"
// @Success		200	{object} []database.Post
// @Failure		404 {object} ErrorMessage
// @Router		/bytag [post]
func getPostByTag(c *gin.Context) {
	var tagObj database.TagObj
	if err := c.BindJSON(&tagObj); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var posts []database.Post
	posts, err := database.GetPostByTag(tagObj.Tag)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
