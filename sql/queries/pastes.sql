-- name: GetPaste :one
SELECT * FROM pastes WHERE id = ? LIMIT 1;
