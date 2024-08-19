-- name: AddUser :execresult
INSERT INTO users (id, email, password, oauth_id, name, role)
VALUES (?, ?, ?, ?, ?, ?);

-- name: EmailExists :one
SELECT COUNT(1) FROM users WHERE email = ?;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = ?;