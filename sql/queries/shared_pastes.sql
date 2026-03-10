-- name: GetSharedPaste :one
SELECT * FROM shared_pastes WHERE id = ? LIMIT 1;
