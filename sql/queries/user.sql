-- name: CreateUser :one
-- Creates a new user and returns their ID.
INSERT INTO users (email, password_hash) VALUES(?, ?) RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;
