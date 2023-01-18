package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListPostsHandler(t *testing.T) {

	// Arrange
	expected := posts[0]

	// Act
	r := SetUpRouter()
	r.GET("/", listPostsHandler)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic("COULDN'T MAKE A REQUEST")
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Incorrect response body.")
	}
	w.Result().Body.Close()
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		t.Fatal(err)
	}
	result := posts[0]

	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, result)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Incorrect response status. Expected %v, got %v", http.StatusOK, w.Code)
	}
}

func TestPostHandler(t *testing.T) {
	expected := posts[0]

	r := SetUpRouter()
	r.GET("/:id", postHandler)
	req, err := http.NewRequest("GET", "/1", nil)
	if err != nil {
		panic("COULDN'T MAKE A REQUEST")
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Incorrect response body.")
	}
	w.Result().Body.Close()
	var result Post
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, result)
	}
}

func TestCreatePostHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/:id", createPostHandler)
	post := Post{
		ID:        2,
		AuthorID:  1,
		Title:     "someTitleText",
		Anons:     "anonsText",
		CreatedAt: "2023-01-17 18:45:25+05:00",
		UpdatedAt: "2023-01-17 18:45:25+05:00",
		Content:   "contentText",
	}
	jsonValue, _ := json.Marshal(post)
	reqFound, err := http.NewRequest("POST", "/2", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic("COULDN'T MAKE A REQUEST")
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	if w.Code != http.StatusOK {
		t.Errorf("Incorrect response status. Expected %v, got %v", http.StatusOK, w.Code)
	}
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Incorrect response body.")
	}
	w.Result().Body.Close()
	var id int64
	json.Unmarshal(body, &id)

	if id != post.ID {
		t.Errorf("Incorrect result. Expected %v, got %v", post.ID, id)
	}
}

func TestEditPostHandler(t *testing.T) {

	post := Post{ID: 1, AuthorID: 3, Title: "", Anons: "", CreatedAt: "", UpdatedAt: "", Content: ""}
	jsonValue, _ := json.Marshal(post)

	r := SetUpRouter()
	r.POST("/:id", editPostHandler)

	reqFound, _ := http.NewRequest("POST", "/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	if w.Code != http.StatusOK {
		t.Errorf("Incorrect response code. Expected %v, got %v", http.StatusOK, w.Code)
	}
}

func TestDeleteHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/:id", deletePostHandler)
	req, err := http.NewRequest("DELETE", "/1", nil)
	if err != nil {
		panic("COULDN'T MAKE A REQUEST!")
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Incorrect response code. Expected %v, got %v.", http.StatusNoContent, w.Code)
	}
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
