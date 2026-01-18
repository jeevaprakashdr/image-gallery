-- name: ListImages :many
SELECT * FROM images;


-- name: GetImage :one
SELECT * FROM images
WHERE title = $1 LIMIT 1;

-- name: SaveImage :one
INSERT INTO images (id , title, tags)
VALUES ($1, $2, $3)
RETURNING *;