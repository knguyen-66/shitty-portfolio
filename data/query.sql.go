// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package data

import (
	"context"
	"database/sql"
)

const insertBlog = `-- name: InsertBlog :one
INSERT INTO blog (title, slug, summary, thumbnail_dir, content_filename, time_created) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, title, slug, summary, thumbnail_dir, content_filename, time_created
`

type InsertBlogParams struct {
	Title           string
	Slug            string
	Summary         sql.NullString
	ThumbnailDir    sql.NullString
	ContentFilename string
	TimeCreated     sql.NullTime
}

func (q *Queries) InsertBlog(ctx context.Context, arg InsertBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, insertBlog,
		arg.Title,
		arg.Slug,
		arg.Summary,
		arg.ThumbnailDir,
		arg.ContentFilename,
		arg.TimeCreated,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Summary,
		&i.ThumbnailDir,
		&i.ContentFilename,
		&i.TimeCreated,
	)
	return i, err
}

const insertBlogTag = `-- name: InsertBlogTag :one
INSERT INTO blog_tag (id_blog, tag_name) VALUES (?, ?) RETURNING id, id_blog, tag_name
`

type InsertBlogTagParams struct {
	IDBlog  int64
	TagName string
}

func (q *Queries) InsertBlogTag(ctx context.Context, arg InsertBlogTagParams) (BlogTag, error) {
	row := q.db.QueryRowContext(ctx, insertBlogTag, arg.IDBlog, arg.TagName)
	var i BlogTag
	err := row.Scan(&i.ID, &i.IDBlog, &i.TagName)
	return i, err
}

const selectBlog = `-- name: SelectBlog :one
SELECT id, title, slug, summary, thumbnail_dir, content_filename, time_created FROM blog WHERE id = ? LIMIT 1
`

func (q *Queries) SelectBlog(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, selectBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Summary,
		&i.ThumbnailDir,
		&i.ContentFilename,
		&i.TimeCreated,
	)
	return i, err
}

const selectBlogBySlug = `-- name: SelectBlogBySlug :one
SELECT id, title, slug, summary, thumbnail_dir, content_filename, time_created FROM blog WHERE slug = ? LIMIT 1
`

func (q *Queries) SelectBlogBySlug(ctx context.Context, slug string) (Blog, error) {
	row := q.db.QueryRowContext(ctx, selectBlogBySlug, slug)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Summary,
		&i.ThumbnailDir,
		&i.ContentFilename,
		&i.TimeCreated,
	)
	return i, err
}

const selectBlogTagsByBlog = `-- name: SelectBlogTagsByBlog :many
SELECT id, id_blog, tag_name FROM blog_tag WHERE id_blog = ?
`

func (q *Queries) SelectBlogTagsByBlog(ctx context.Context, idBlog int64) ([]BlogTag, error) {
	rows, err := q.db.QueryContext(ctx, selectBlogTagsByBlog, idBlog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BlogTag
	for rows.Next() {
		var i BlogTag
		if err := rows.Scan(&i.ID, &i.IDBlog, &i.TagName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectBlogTagsWithOccurance = `-- name: SelectBlogTagsWithOccurance :many
SELECT id, id_blog, tag_name, count(1) as occurences FROM blog_tag GROUP BY tag_name ORDER BY occurences DESC
`

type SelectBlogTagsWithOccuranceRow struct {
	ID         int64
	IDBlog     int64
	TagName    string
	Occurences int64
}

func (q *Queries) SelectBlogTagsWithOccurance(ctx context.Context) ([]SelectBlogTagsWithOccuranceRow, error) {
	rows, err := q.db.QueryContext(ctx, selectBlogTagsWithOccurance)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectBlogTagsWithOccuranceRow
	for rows.Next() {
		var i SelectBlogTagsWithOccuranceRow
		if err := rows.Scan(
			&i.ID,
			&i.IDBlog,
			&i.TagName,
			&i.Occurences,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectBlogs = `-- name: SelectBlogs :many
SELECT id, title, slug, summary, thumbnail_dir, content_filename, time_created FROM blog ORDER BY time_created
`

func (q *Queries) SelectBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, selectBlogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Slug,
			&i.Summary,
			&i.ThumbnailDir,
			&i.ContentFilename,
			&i.TimeCreated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
