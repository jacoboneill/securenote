-- name: CreateUser :one
INSERT INTO users (email, password_hash) VALUES(?, ?) RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;
