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

-- name: UpdateTitle :one
UPDATE articles
set title = $2
WHERE article_id = $1
RETURNING *;

-- name: UpdateArticleContent :one
UPDATE articles
set content = $2
WHERE article_id = $1
RETURNING *;

-- name: UpdateArticleChangeAt :one
UPDATE articles
set changed_at = $2
WHERE article_id = $1
RETURNING *;
