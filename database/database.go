package database

import (
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

type Dbinstance struct {
	DB *sqlx.DB
}

var Postgres Dbinstance

func GetAll() ([]Post, error) {
	posts := []Post{}

	db, err := sqlx.Connect("postgres", "user=postgres password=3729 dbname=blogpost sslmode=disable") //
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("DB connected")

	rows, err := db.Query(`SELECT * FROM tbl_post;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetList(num int) ([]Post, error) {
	posts := []Post{}

	rows, err := Postgres.DB.Query(`SELECT * FROM tbl_post limit %d;`, num)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		if err != nil {
			return nil, err
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

	log.Println("DB connected")

	Postgres = Dbinstance{
		DB: db,
	}
}
