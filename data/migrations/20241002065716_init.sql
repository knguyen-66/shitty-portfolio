-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blog (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
  	title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
  	summary TEXT,
  	thumbnail_dir TEXT,
  	content_filename TEXT NOT NULL,
    time_created DATETIME
);

CREATE TABLE IF NOT EXISTS blog_tag (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  id_blog INTEGER NOT NULL,
  tag_name TEXT NOT NULL,
  FOREIGN KEY(id_blog) REFERENCES blog(id) ON UPDATE CASCADE
);

CREATE INDEX idx_blog__slug ON blog(slug);
CREATE INDEX idx_blog_tag__id_blog ON blog_tag(id_blog);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blog_tag;
DROP TABLE IF EXISTS blog;
-- +goose StatementEnd
