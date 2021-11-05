-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (
  name, age
) VALUES (
  ?, ?
);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;