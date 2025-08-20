package databse

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/gowaliullah/blog-rest-api/models"
// 	_ "github.com/lib/pq"
// )

// // PostgresDB wraps the SQL database connection
// type PostgresDB struct {
// 	*sql.DB
// }

// // NewPostgresDB initializes a new PostgreSQL connection
// func DB(cfg *models.Config) (*PostgresDB, error) {
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database: %w", err)
// 	}

// 	return &PostgresDB{db}, nil
// }
