package api

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GenerateOtpDto struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func GenerateOtp(ctx *gin.Context, q *db.Queries, _ *pgx.Conn) {

	var generateOtpDto GenerateOtpDto
	if err := ctx.ShouldBindBodyWithJSON(&generateOtpDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := q.GetUserByPhoneNumber(ctx, pgtype.Text{Valid: true, String: generateOtpDto.PhoneNumber})

	if err != nil {
		// the user is found
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	otp := fmt.Sprintf("%04d", rand.Intn(10000))
	timeAfterMin := time.Now().UTC().Add(time.Minute)

	updatedUser, err := q.UpdateOtp(ctx, db.UpdateOtpParams{
		Otp:               pgtype.Text{String: otp, Valid: true},
		OtpExpirationTime: pgtype.Timestamp{Time: timeAfterMin, Valid: true},
		PhoneNumber:       pgtype.Text{String: generateOtpDto.PhoneNumber, Valid: true},
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
	})

}
