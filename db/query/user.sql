-- name: GetUserByID :one
SELECT id, name, email 
FROM users 
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email) 
VALUES ($1, $2) 
RETURNING id, name, email;

-- name: ListUsers :many
SELECT id, name, email 
FROM users;
-- name: updateUser :one
UPDATE users 
SET name = $2, email = $3 
WHERE id = $1 
RETURNING id, name, email;