-- name: AddTransaction :execresult
INSERT INTO transactions (id, user_id, descript)
VALUES (?, ?, ?);

-- name: GetAllByUserID :many
SELECT * FROM transactions WHERE user_id = ?;