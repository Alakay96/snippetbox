package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"snippetbox.alex.net/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// .env file is optional, so we don't exit if it's not found
		fmt.Println("Warning: .env file not found, using defaults or command line flags")
	}

	// Get default values from environment variables
	defaultAddr := os.Getenv("APP_PORT")
	if defaultAddr == "" {
		defaultAddr = "4000"
	}
	defaultAddr = ":" + defaultAddr

	defaultDSN := buildDSN()

	addr := flag.String("addr", defaultAddr, "HTTP network address")
	dsn := flag.String("dsn", defaultDSN, "MariaDB data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// buildDSN constructs the database connection string from environment variables
func buildDSN() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "web"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "pass"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "snippetbox"
	}

	charset := os.Getenv("DB_CHARSET")
	if charset == "" {
		charset = "utf8mb4"
	}

	parseTime := os.Getenv("DB_PARSE_TIME")
	if parseTime == "" {
		parseTime = "true"
	}

	// Build the DSN string
	if port != "3306" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
			user, password, host, port, dbname, charset, parseTime)
	}

	// For standard port, use the shorter format
	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s",
		user, password, dbname, charset, parseTime)
}
