package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"testrepo/database"

	"github.com/gin-gonic/gin"
)

type Router struct {
	R *gin.Engine
}

var Tool Router

func TestMain(m *testing.M) {
	database.ConnectDB()
	Tool.SetUpRouter()
	Tool.R.GET("/list", allPostsHandler)
	Tool.R.GET("/list/:num", listPostsHandler)
	Tool.R.GET("/post/:id", postHandler)
	Tool.R.POST("/post", createPostHandler)
	Tool.R.DELETE("/post/:id", deletePostHandler)
	Tool.R.POST("/post/:id", editPostHandler)
	Tool.R.POST("/bytag", getPostByTag)
	code := m.Run()
	database.Postgres.DB.Close()
	os.Exit(code)
}

func (tool *Router) SetUpRouter() {
	tool.R = gin.Default()
}

func TestAllPostHandler(t *testing.T) {

	// Gather expected information
	rows, err := database.Postgres.DB.Query(`SELECT * FROM tbl_post`)
	if err != nil {
		t.Errorf("Error while query database: %v", err.Error())
	}
	expected := []database.Post{}
	for rows.Next() {
		var post database.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID); err != nil {
			t.Errorf("Error while scanning row: %v", err.Error())
		}
		expected = append(expected, post)
	}

	// Make request, save response into varible "got"
	req, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		t.Errorf("Error while making new request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)

	responseData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v", err.Error())
	}
	got := []database.Post{}

	if err = json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("Error while unmarshaling responseData: %v", err.Error())
	}

	// Check if result not correct, print error
	if !SlicesEqual(expected, got) {
		t.Errorf("ERROR, expacted %v, got %v", expected, got)
	}
}

func TestListPostsHandler(t *testing.T) {

	// Gather expected information
	i := 50 // option
	// for i := 1; i < 500; i++ {
	expected, err := database.GetList(i)
	if err != nil {
		t.Errorf("Error while query. %v", err.Error())
	}

	// Make request, save response into "got" varible
	req, err := http.NewRequest("GET", fmt.Sprintf("/list/%d", i), nil)
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: %v", fmt.Sprintf("/list/%d", i))
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)
	responseData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v, iteration: %d", err.Error(), i)
	}
	got := []database.Post{}
	if err = json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("Error while unmarshaling responseData: %v", err.Error())
	}

	// Check if result not correct, print error
	if !SlicesEqual(expected, got) {
		t.Errorf("ERROR, expacted %v, got %v", expected, got)
	}
	// }
}

func TestPostHandler(t *testing.T) {

	// for i := 1; i < 500; i++ {
	i := 5
	// Gather expected information
	expected, err := database.GetPost(int64(i))
	if err != nil {
		t.Errorf("Error while getting post from DB: %v", err)
	}

	// Make request, save response into "got" varible
	req, err := http.NewRequest("GET", fmt.Sprintf("/post/%d", i), nil)
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: %v", "/post/1")
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)
	responseData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v.", err.Error())
	}
	got := database.Post{}
	if err = json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("Error while unmarshaling responseData: %v", err.Error())
	}

	// Check if result not correct, print error
	if expected != got {
		t.Errorf("Error, expected: %v, got: %v.", expected, got)
	}
	// }
}

func TestCreatePostHandler(t *testing.T) {

	// Make request for create post
	var jsonStr = []byte(`{"id":"1","title":"СГБ обнаружила очередные тоннели, связывающие Узбекистан и Кыргызстан — видео", "content":"B Кургантепинском районе Андижанской области обнаружены подземные тоннели, ведущие в Кыргызстан. Сообщается, что 3 и 9 февраля в ходе оперативных мероприятий, проводимых СГБ и МВД, были задержаны лица, незаконно проникшие на территорию Кургантепинского района Андижанской области территории -Сууского района Ошской области Кыргызстана.","author_id":"1"}`)
	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: %v", "/post")
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)

	// Check code result if OK
	if w.Code != http.StatusOK {
		t.Errorf("Error, expected: %v, got: %v.", http.StatusOK, w.Code)
	}

	// Delete created post from DB
	list, err := database.GetAll()
	if err != nil {
		t.Errorf("Error while query database: %v", err.Error())
	}
	for _, v := range list {
		if v.Title == "СГБ обнаружила очередные тоннели, связывающие Узбекистан и Кыргызстан — видео" {
			if err := database.DeletePost(int(v.ID)); err != nil {
				log.Printf("Cannot delete created post, ID: %d. Delete this manualy", v.ID)
			}
		}
	}
}

func TestDeletePostHandler(t *testing.T) {

	// Adding new post to DB for emulating
	id, err := database.AddPost(database.Post{ID: 1, Title: "СГБ обнаружила очередные тоннели", Content: "СГБ обнаружила очередные тоннели, связывающие Узбекистан и Кыргызстан — видео", AuthorID: 1})
	if err != nil {
		t.Error("Error while emulating post, cannot add row to DB.")
	}

	// Making request for delete emulated post
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/post/%d", id), nil)
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: /post/%d", id)
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)

	// Check if status code, if OK, then created post is alredy deleted when request
	if w.Code != http.StatusOK {
		t.Errorf("Error, expected: %v, got: %v.", http.StatusOK, w.Code)
	}
}

func TestEditPostHandler(t *testing.T) {

	// Adding new post to DB for emulating
	id, err := database.AddPost(database.Post{ID: 1, Title: "СГБ обнаружила очередные тоннели", Content: "СГБ обнаружила очередные тоннели, связывающие Узбекистан и Кыргызстан — видео", AuthorID: 1})
	if err != nil {
		t.Error("Error while emulating post, cannot add row to DB.")
	}

	// Make request for edit post from DB
	var jsonStr = []byte(fmt.Sprintf(`{"id":"%d","title":"СГБ обнаружила очередные тоннели", "content":"B Кургантепинском районе Андижанской области обнаружены подземные тоннели.","author_id":"2"}`, id))
	req, err := http.NewRequest("POST", fmt.Sprintf("/post/%d", id), bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: /post/%d", id)
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Error, expected: %v, got: %v.", http.StatusOK, w.Code)
	}

	// Delete emulated post from DB
	if err = database.DeletePost(id); err != nil {
		log.Printf("Cannot delete created post, ID: %d. Delete this manualy", id)
	}
}

func TestGetPostByTag(t *testing.T) {

	// Making request and getting data
	jsonString := []byte(`{"tag": "#энергетика"}`)
	req, err := http.NewRequest("POST", "/bytag", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Errorf("Error while making new request. ReqUrl: /bytag")
	}
	w := httptest.NewRecorder()
	Tool.R.ServeHTTP(w, req)
	responseData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error while reading response body: %v.", err.Error())
	}
	got := []database.Post{}
	if err = json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("Error while unmarshaling responseData: %v", err.Error())
	}

	// Gather expected information from DB
	expected, err := database.GetPostByTag("#энергетика")
	if err != nil {
		t.Error("Error when getting data from DB by tag #энергетика")
	}

	// Check goted information and expected data
	if !SlicesEqual(expected, got) {
		t.Errorf("ERROR expected: %v, got: %v.", expected, got)
	}
}

func SlicesEqual(a, b []database.Post) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
