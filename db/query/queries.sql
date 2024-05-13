-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, password, phone_number, otp, otp_expiration_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateProfile :one
INSERT INTO profile (user_id, first_name, last_name, address)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateOtp :one
UPDATE users SET otp = $1, otp_expiration_time = $2 where phone_number = $3 RETURNING *;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: GetUsers :many
SELECT * FROM users u JOIN profile p on u.id = p.user_id OFFSET $1 LIMIT $2;