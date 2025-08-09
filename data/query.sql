-- name: SelectBlog :one
SELECT * FROM blog WHERE id = ? LIMIT 1;

-- name: SelectBlogByContentFilename :one
SELECT * FROM blog WHERE content_filename = ? LIMIT 1;

-- name: SelectBlogBySlug :one
SELECT * FROM blog WHERE slug = ? LIMIT 1;

-- name: SelectBlogs :many
SELECT * FROM blog ORDER BY time_created;

-- name: InsertBlog :one
INSERT INTO blog (title, slug, summary, thumbnail_dir, content_filename, time_created) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: SelectBlogTagsWithOccurance :many
SELECT *, count(1) as occurences FROM blog_tag GROUP BY tag_name ORDER BY occurences DESC;

-- name: SelectBlogTagsByBlog :many
SELECT * FROM blog_tag WHERE id_blog = ?;

-- name: InsertBlogTag :one
INSERT INTO blog_tag (id_blog, tag_name) VALUES (?, ?) RETURNING *;
