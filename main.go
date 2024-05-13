package main

import (
	"context"
	"fmt"
	"os"

	"github.com/MohammadAzhari/aqary_task/api"
	"github.com/MohammadAzhari/aqary_task/db"
)

func main() {
	conn, err := db.NewConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	q := db.New(conn)

	api.NewServer(q, conn)
}
