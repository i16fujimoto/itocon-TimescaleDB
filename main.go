package main

import (
   "context"
   "fmt"
   "os"
   "time"

   //"github.com/jackc/pgx/v4"
   "github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	ctx := context.Background()
	connStr := os.Getenv("CONNECT")
	dbpool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	//Execute query on TimescaleDB
	rows, err := dbpool.Query(ctx, "select * from sensor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute query %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	fmt.Println("Successfully executed query", rows)

	type result struct {
		sensor_id string
		sensor_name string
		user_id string
		is_active bool
		created_at time.Time
	}
	for rows.Next() {
		var r result
		err = rows.Scan(&r.sensor_id, &r.sensor_name, &r.user_id, &r.is_active, &r.created_at)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
			os.Exit(1)
		}
		fmt.Println(&r.sensor_id, &r.sensor_name, &r.user_id, &r.is_active, &r.created_at)
	}



}