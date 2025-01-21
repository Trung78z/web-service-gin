package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(dsn string) {
	// Sử dụng context.Background() để tạo kết nối
	ctx := context.Background()

	// Khởi tạo kết nối với pool cơ sở dữ liệu
	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Database connection initialized")
}
