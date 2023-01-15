// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: article.sql

package db

import (
	"context"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (
  author_username,
  title,
  content
) VALUES (
  $1, $2, $3
) RETURNING article_id, title, author_username, content
`

type CreateArticleParams struct {
	AuthorUsername string `json:"author_username"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, createArticle, arg.AuthorUsername, arg.Title, arg.Content)
	var i Article
	err := row.Scan(
		&i.ArticleID,
		&i.Title,
		&i.AuthorUsername,
		&i.Content,
	)
	return i, err
}

const deleteArticle = `-- name: DeleteArticle :exec
DELETE FROM articles
WHERE article_id = $1
`

func (q *Queries) DeleteArticle(ctx context.Context, articleID int64) error {
	_, err := q.db.ExecContext(ctx, deleteArticle, articleID)
	return err
}