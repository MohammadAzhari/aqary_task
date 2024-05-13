package api

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"

	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetUsers(ctx *gin.Context, q *db.Queries, conn *pgx.Conn) {
	// count the users
	cnt, err := q.CountUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pageSize := int32(2)
	numOfGoRoutines := int32(math.Ceil(float64(cnt) / float64(pageSize)))

	usersChan := make(chan []db.GetUsersRow, numOfGoRoutines)

	var wg sync.WaitGroup
	wg.Add(int(numOfGoRoutines))
	// fire threads to scan the users
	for i := int32(0); i < numOfGoRoutines; i++ {
		go func() {
			defer wg.Done()
			offset := i * pageSize
			limit := pageSize

			conn, err = db.NewConn()
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

			defer conn.Close(ctx)
			q := db.New(conn)

			res, err := q.GetUsers(ctx, db.GetUsersParams{Offset: offset, Limit: limit})
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
			usersChan <- res
		}()
	}

	go func() {
		wg.Wait()
		close(usersChan)
	}()

	users := make([]db.GetUsersRow, 0)

	for u := range usersChan {
		users = append(users, u...)
	}

	// sort the users
	sort.Slice(users, func(i, j int) bool {
		return users[i].FirstName.String < users[j].FirstName.String
	})

	ctx.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}
