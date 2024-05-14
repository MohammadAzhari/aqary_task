package db_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/MohammadAzhari/aqary_task/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func Test_Tx_Success(t *testing.T) {
	conn, err := db.InitPool()

	if err != nil {
		t.Error(err)
	}

	q := db.New(conn)

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		t.Error(err)
	}

	q = q.WithTx(tx)
	defer tx.Rollback(context.Background())

	user, profile, err := createRandomUser(q)

	if err != nil {
		t.Error(err)
	}

	if user.ID != profile.UserID.Int32 {
		t.Errorf("user id not the same as the profile userID")
	}
}

func Test_Tx_Fail(t *testing.T) {
	conn, err := db.InitPool()

	if err != nil {
		t.Error(err)
	}

	q := db.New(conn)

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	q = q.WithTx(tx)
	defer tx.Rollback(context.Background())

	user, _, err := createRandomUser(q)

	if err != nil {
		t.Error(err)
	}

	if user != nil {
		_, err = q.CreateUser(context.Background(), db.CreateUserParams{
			PhoneNumber: user.PhoneNumber,
		})
	}

	// doublicate the user:
	if err == nil {
		t.Errorf("the doublicate doesn't works")
	}
}

func createRandomUser(q *db.Queries) (*db.User, *db.Profile, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	PhoneNumber := fmt.Sprintf("%04d", rand.Intn(10000))

	user, err := q.CreateUser(context.Background(), db.CreateUserParams{
		PhoneNumber: pgtype.Text{
			String: PhoneNumber,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, nil, err
	}

	profile, err := q.CreateProfile(context.Background(), db.CreateProfileParams{
		FirstName: pgtype.Text{
			String: "john",
			Valid:  true,
		},
		UserID: pgtype.Int4{
			Int32: user.ID,
			Valid: true,
		},
	})

	return &user, &profile, err
}
