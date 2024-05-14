package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateUserDto struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func CreateUser(ctx *gin.Context, q *db.Queries, conn *pgxpool.Pool) {

	var userDto CreateUserDto
	if err := ctx.ShouldBindBodyWithJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := q.GetUserByPhoneNumber(ctx, pgtype.Text{Valid: true, String: userDto.PhoneNumber})

	if err == nil {
		// the user is found
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User is already found"})
		return
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		// if the error is not ErrNoRows
		fmt.Println(err, pgx.ErrNoRows)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		// if the error is not ErrNoRows
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	qtx := q.WithTx(tx)

	user, err := qtx.CreateUser(ctx, db.CreateUserParams{
		PhoneNumber: pgtype.Text{String: userDto.PhoneNumber, Valid: true},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile, err := qtx.CreateProfile(ctx, db.CreateProfileParams{
		UserID:    pgtype.Int4{Int32: user.ID, Valid: true},
		FirstName: pgtype.Text{String: userDto.Name, Valid: true},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit(ctx)

	ctx.JSON(http.StatusCreated, gin.H{
		"profile": profile,
		"user":    user,
	})

}
