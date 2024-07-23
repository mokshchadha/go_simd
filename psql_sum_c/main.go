package main

/*
#include "simd_sum.c"
*/
import "C"
import (
	"context"
	"fmt"
	"log"
	"time"
	"unsafe"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connectToDB(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return pool, nil
}

func fetchRecords(pool *pgxpool.Pool, limit int) ([]int32, error) {
	query := fmt.Sprintf("SELECT startYear FROM staging_titles st LIMIT %d", limit)
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	results := []int32{}

	for rows.Next() {
		var startYear *int32
		err := rows.Scan(&startYear)
		if err != nil {
			return nil, fmt.Errorf("failed to get row values: %v", err)
		}

		if startYear != nil {
			results = append(results, *startYear)
		}
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, nil
}

func calculateSum(records []int32) int32 {
	var sum int32
	for _, value := range records {
		sum += value
	}
	return sum
}

func calculateSumSIMD(records []int32) int32 {
	length := len(records)
	return int32(C.simd_sum((*C.int)(unsafe.Pointer(&records[0])), C.int(length)))
}

func measureExecutionTime(pool *pgxpool.Pool, limit int) {
	records, err := fetchRecords(pool, limit)
	if err != nil {
		log.Fatalf("Failed to fetch records: %v", err)
	}
	start := time.Now()

	sum := calculateSum(records)
	duration := time.Since(start)

	fmt.Printf("Fetched %d records and sum of %d in %s\n", limit, sum, duration)
}

func measureExecutionTimeSIMD(pool *pgxpool.Pool, limit int) {
	records, err := fetchRecords(pool, limit)
	if err != nil {
		log.Fatalf("Failed to fetch records: %v", err)
	}
	start := time.Now()

	sum := calculateSumSIMD(records)
	duration := time.Since(start)

	fmt.Printf("Fetched %d records and SIMD sum of %d in %s\n", limit, sum, duration)
}

func main() {
	connString := "postgres://moksh@localhost/postgres"

	pool, err := connectToDB(connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	fmt.Println("Successfully connected to the database!")

	measureExecutionTime(pool, 512)
	measureExecutionTime(pool, 1024)
	measureExecutionTime(pool, 10240)
	measureExecutionTime(pool, 102400)
	measureExecutionTime(pool, 1024000)

	fmt.Println("=======================")

	measureExecutionTimeSIMD(pool, 512)
	measureExecutionTimeSIMD(pool, 1024)
	measureExecutionTimeSIMD(pool, 10240)
	measureExecutionTimeSIMD(pool, 102400)
	measureExecutionTimeSIMD(pool, 1024000)
}
