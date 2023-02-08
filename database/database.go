package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}

type TagObj struct {
	Tag string `json:"tag"`
}

type Dbinstance struct {
	DB *sqlx.DB
}

var Postgres Dbinstance

func GetAll() ([]Post, error) {
	rows, err := Postgres.DB.Query(`SELECT * FROM tbl_post;`)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetList(num int) ([]Post, error) {
	rows, err := Postgres.DB.Query(fmt.Sprintf(`SELECT * FROM tbl_post limit %d`, num))
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPost(id int64) (Post, error) {
	var post Post
	rows, err := Postgres.DB.Query(fmt.Sprintf(`SELECT * FROM tbl_post WHERE id=%d LIMIT 1`, id))
	if err != nil {
		return Post{}, err
	}

	ok := rows.Next()
	if ok {
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID); err != nil {
			return Post{}, err
		}
	}
	return post, nil
}

func AddPost(newpost Post) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO tbl_post (title, content, author_id) VALUES ('%v', '%v', %d) RETURNING id`, newpost.Title, newpost.Content, newpost.AuthorID)
	if err := Postgres.DB.QueryRow(query).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func DeletePost(id int) error {
	if _, err := Postgres.DB.Exec(`DELETE FROM tbl_post WHERE id=$1`, id); err != nil {
		return err
	}
	return nil
}

func EditPost(post Post) error {
	if _, err := Postgres.DB.Exec(`UPDATE tbl_post SET title=$1, content=$2, author_id=$3 WHERE id=$4`, post.Title, post.Content, post.AuthorID, post.ID); err != nil {
		return err
	}
	return nil
}

func GetPostByTag(tag string) ([]Post, error) {
	rows, err := Postgres.DB.Query(`select * from tbl_post where id in(select post_id from post_tag where tag_id in(select id from tbl_tag where name=$1))`, tag)
	if err != nil {
		return []Post{}, err
	}

	var posts []Post
	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID); err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func ConnectDB() {
	db, err := sqlx.Connect("postgres", "user=postgres password=3729 dbname=blogpost sslmode=disable") //
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("DB connect success")

	Postgres = Dbinstance{
		DB: db,
	}
}
