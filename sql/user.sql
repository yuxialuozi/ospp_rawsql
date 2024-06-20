-- name: InsertOne :execresult
INSERT INTO users (username, email, age) VALUES (?, ?, ?);
-- name: InsertMany :execresult
INSERT INTO users (username, email, age) VALUES (?, ?, ?);
-- name: FindUsernames :many
SELECT * FROM users ORDER BY id;
-- name: FindByUsernameAge :one
SELECT * FROM users WHERE id = ? LIMIT 1;
-- name: UpdateContact :exec
DELETE FROM users WHERE id = ?;
-- name: DeleteById :exec
DELETE FROM users WHERE id = ?;
-- name: CountByAge :one
SELECT * FROM users WHERE id = ? LIMIT 1;