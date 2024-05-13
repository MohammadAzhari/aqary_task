// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProfile = `-- name: CreateProfile :one
INSERT INTO profile (user_id, first_name, last_name, address)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, first_name, last_name, address
`

type CreateProfileParams struct {
	UserID    pgtype.Int4
	FirstName pgtype.Text
	LastName  pgtype.Text
	Address   pgtype.Text
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, createProfile,
		arg.UserID,
		arg.FirstName,
		arg.LastName,
		arg.Address,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Address,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password, phone_number, otp, otp_expiration_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, password, phone_number, otp, otp_expiration_time
`

type CreateUserParams struct {
	Email             pgtype.Text
	Password          pgtype.Text
	PhoneNumber       pgtype.Text
	Otp               pgtype.Text
	OtpExpirationTime pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Password,
		arg.PhoneNumber,
		arg.Otp,
		arg.OtpExpirationTime,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const getUserByPhoneNumber = `-- name: GetUserByPhoneNumber :one
SELECT id, email, password, phone_number, otp, otp_expiration_time FROM users WHERE phone_number = $1 LIMIT 1
`

func (q *Queries) GetUserByPhoneNumber(ctx context.Context, phoneNumber pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhoneNumber, phoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT u.id, email, password, phone_number, otp, otp_expiration_time, p.id, user_id, first_name, last_name, address FROM users u JOIN profile p on u.id = p.user_id OFFSET $1 LIMIT $2
`

type GetUsersParams struct {
	Offset int32
	Limit  int32
}

type GetUsersRow struct {
	ID                int32
	Email             pgtype.Text
	Password          pgtype.Text
	PhoneNumber       pgtype.Text
	Otp               pgtype.Text
	OtpExpirationTime pgtype.Timestamp
	ID_2              int32
	UserID            pgtype.Int4
	FirstName         pgtype.Text
	LastName          pgtype.Text
	Address           pgtype.Text
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.PhoneNumber,
			&i.Otp,
			&i.OtpExpirationTime,
			&i.ID_2,
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOtp = `-- name: UpdateOtp :one
UPDATE users SET otp = $1, otp_expiration_time = $2 where phone_number = $3 RETURNING id, email, password, phone_number, otp, otp_expiration_time
`

type UpdateOtpParams struct {
	Otp               pgtype.Text
	OtpExpirationTime pgtype.Timestamp
	PhoneNumber       pgtype.Text
}

func (q *Queries) UpdateOtp(ctx context.Context, arg UpdateOtpParams) (User, error) {
	row := q.db.QueryRow(ctx, updateOtp, arg.Otp, arg.OtpExpirationTime, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}