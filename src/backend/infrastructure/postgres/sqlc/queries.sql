-- name: ListImages :many
SELECT * FROM images;


-- name: GetImage :one
SELECT * FROM images
WHERE title = $1 LIMIT 1;