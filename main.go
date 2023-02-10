package main

// planning				DONE
// Development			DONE
// UnitTesting			DONE
// DocumentationSwaggo	PROGRESS
// Make Public in Git	PROGRESS

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testrepo/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	database.ConnectDB()

	r.GET("/list", allPostsHandler)       //GET ALL LIST
	r.GET("/list/:num", listPostsHandler) //GET LIST

	r.GET("/post/:id", postHandler)          //GET POST
	r.POST("/post", createPostHandler)       //CREATE POST
	r.DELETE("/post/:id", deletePostHandler) //DELETE POST
	r.POST("/post/:id", editPostHandler)     //EDIT POST

	r.POST("/bytag", getPostByTag) //POST POST BY TAG

	r.Run(":9000")
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

func allPostsHandler(c *gin.Context) {
	posts, err := database.GetAll()
	if err != nil {
		jsonValue, _ := json.Marshal(err)
		c.JSON(http.StatusInternalServerError, jsonValue)
		return
	}
	c.JSON(http.StatusOK, posts)
}

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
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

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

func deletePostHandler(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err = database.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

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
