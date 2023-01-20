--- psql -U postgres -d blogpost -a -f Desktop/insert_blogs.sql
CREATE TABLE IF NOT EXISTS tbl_user (
	id SERIAL PRIMARY KEY,
	username VARCHAR(25) NOT NULL,
	password VARCHAR(25) NOT NULL,
	email VARCHAR(50) NOT NULL,
	profile TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS tbl_post (
	id SERIAL PRIMARY KEY,
	title VARCHAR(250) NOT NULL,
	content TEXT NOT NULL,
	author_id INT NOT NULL,
	FOREIGN KEY (author_id) REFERENCES tbl_user(id));

CREATE TABLE IF NOT EXISTS tbl_comment (
	id SERIAL PRIMARY KEY,
	content TEXT NOT NULL,
	create_time timestamptz NOT NULL,
	author_id INT NOT NULL,
	post_id INT NOT NULL,
	FOREIGN KEY (post_id) REFERENCES tbl_post(id),
	FOREIGN KEY (author_id) REFERENCES tbl_user(id));

CREATE TABLE IF NOT EXISTS tbl_tag (
	id SERIAL PRIMARY KEY,
	name VARCHAR(25));

CREATE TABLE IF NOT EXISTS post_tag (
	post_id INT NOT NULL,
	tag_id INT NOT NULL,
	FOREIGN KEY (post_id) REFERENCES tbl_post(id),
	FOREIGN KEY (tag_id) REFERENCES tbl_tag(id));
