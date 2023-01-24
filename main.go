package main

import (
	"encoding/json"
	"net/http"
	"testrepo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	database.ConnectDB()

	r.GET("/list", allPostsHandler) //GET ALL LIST
	// r.GET("/list/:num", listPostsHandler) //GET LIST

	// r.GET("/post/:id", postHandler)          //GET POST
	// r.POST("/post/", createPostHandler)      //CREATE POST
	// r.DELETE("/post/:id", deletePostHandler) //DELETE POST
	// r.POST("/post:id", editPostHandler)      //EDIT POST

	// r.GET("/tag/:tag", postWithTag) //GET POST WITH TAG

	r.Run()
}

type Post struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
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
	}
	c.JSON(http.StatusOK, posts)
}

// func listPostsHandler(c *gin.Context) {
// 	// s := c.Param("num")
// 	c.JSON(http.StatusOK, Post{})
// }

// func postHandler(c *gin.Context) {
// 	s := c.Param("id")
// 	n, err := strconv.ParseInt(s, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, a := range posts {
// 		if a.ID == n {
// 			c.JSON(http.StatusOK, posts[i])
// 			return
// 		}
// 	}

// 	c.Status(http.StatusNotFound)
// }

// func createPostHandler(c *gin.Context) {
// 	newpostID := len(posts) + 1

// 	var newpost Post
// 	if err := c.BindJSON(&newpost); err != nil {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}
// 	newpost.ID = int64(newpostID)

// 	posts = append(posts, newpost)
// 	c.JSON(http.StatusOK, newpostID)
// }

// func deletePostHandler(c *gin.Context) {
// 	s := c.Param("id")
// 	n, err := strconv.ParseInt(s, 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i, a := range posts {
// 		if a.ID == n {
// 			posts = append(posts[:i], posts[i+1:]...)
// 		}
// 	}
// 	c.Status(http.StatusNoContent)
// }

// func editPostHandler(c *gin.Context) {

// 	s := c.Param("id")
// 	n, err := strconv.ParseInt(s, 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i, a := range posts {
// 		if a.ID == n {

// 			var post Post
// 			if err := c.ShouldBindJSON(&post); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"error": err.Error(),
// 				})
// 				panic(err)
// 			}

// 			posts = append(posts[:i], posts[i+1:]...)
// 			posts = append(posts, post)
// 			c.Status(http.StatusOK)
// 			return
// 		}
// 	}

// 	c.Status(http.StatusBadRequest)
// }

// func postWithTag(c *gin.Context) {
// }
