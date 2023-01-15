-- name: CreateArticle :one
INSERT INTO articles (
  author_username,
  title,
  content
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE article_id = $1;