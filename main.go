package main

import (
	"fmt"
	"os"

	"github.com/MohammadAzhari/aqary_task/api"
	"github.com/MohammadAzhari/aqary_task/db"
)

func main() {
	pool, err := db.InitPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	q := db.New(pool)

	api.NewServer(q, pool)
}
