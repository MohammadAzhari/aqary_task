package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyOtpDto struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Otp         string `json:"otp" binding:"required"`
}

func VerifyOtp(ctx *gin.Context, q *db.Queries, _ *pgx.Conn) {

	var verifyOtpDto VerifyOtpDto
	if err := ctx.ShouldBindBodyWithJSON(&verifyOtpDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := q.GetUserByPhoneNumber(ctx, pgtype.Text{Valid: true, String: verifyOtpDto.PhoneNumber})

	if err != nil {
		// the user is found
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Otp.String != verifyOtpDto.Otp {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "otp is wrong"})
		return
	}

	if user.OtpExpirationTime.Time.Before(time.Now().UTC()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "otp is expired"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "otp is verified successfully"})
}
