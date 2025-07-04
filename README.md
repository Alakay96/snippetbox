# Snippetbox

A web application for creating and sharing code snippets, built with Go.

## Features

- Create and store code snippets
- View individual snippets
- Clean, responsive web interface
- MySQL/MariaDB database integration

## Environment Configuration

This application uses environment variables for configuration. Create a `.env` file in the project root to set your configuration values.

### Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit the `.env` file with your specific configuration:
   ```bash
   nano .env
   ```

### Environment Variables

#### Database Configuration
- `DB_HOST` - Database host (default: `localhost`)
- `DB_PORT` - Database port (default: `3306`)
- `DB_USER` - Database username (default: `web`)
- `DB_PASSWORD` - Database password (default: `pass`)
- `DB_NAME` - Database name (default: `snippetbox`)
- `DB_CHARSET` - Database charset (default: `utf8mb4`)
- `DB_PARSE_TIME` - Parse time values (default: `true`)

#### Application Configuration
- `APP_PORT` - Application port (default: `4000`)
- `APP_ENV` - Application environment (default: `development`)

#### Security Configuration
- `SESSION_SECRET` - Secret key for session management (generate a secure 32-character string)
- `CSRF_SECRET` - Secret key for CSRF protection (generate a secure 32-character string)

### Example .env File

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_database_username
DB_PASSWORD=your_database_password
DB_NAME=snippetbox
DB_CHARSET=utf8mb4
DB_PARSE_TIME=true

# Application Configuration
APP_PORT=4000
APP_ENV=development

# Security (generate secure random strings for production)
SESSION_SECRET=your-32-character-secret-key-here
CSRF_SECRET=your-32-character-csrf-secret-here
```

## Database Setup

1. Create a MySQL/MariaDB database:
   ```sql
   CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

2. Create a user and grant privileges:
   ```sql
   CREATE USER 'web'@'localhost' IDENTIFIED BY 'your_password';
   GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
   ```

3. Create the snippets table:
   ```sql
   USE snippetbox;
   
   CREATE TABLE snippets (
       id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
       title VARCHAR(100) NOT NULL,
       content TEXT NOT NULL,
       created DATETIME NOT NULL,
       expires DATETIME NOT NULL
   );
   
   CREATE INDEX idx_snippets_created ON snippets(created);
   ```

## Running the Application

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Build the application:
   ```bash
   go build -o bin/web ./cmd/web
   ```

3. Run the application:
   ```bash
   ./bin/web
   ```

   Or run directly with Go:
   ```bash
   go run ./cmd/web
   ```

4. Access the application at `http://localhost:4000`

## Command Line Flags

The application supports command line flags that override environment variables:

- `-addr` - HTTP network address (overrides `APP_PORT`)
- `-dsn` - Database connection string (overrides database environment variables)

Example:
```bash
./bin/web -addr=:8080 -dsn="user:pass@/snippetbox?parseTime=true"
```

## Project Structure

```
snippetbox/
├── cmd/
│   └── web/
│       ├── main.go          # Application entry point
│       ├── handlers.go      # HTTP handlers
│       ├── helpers.go       # Helper functions
│       └── routes.go        # Route definitions
├── internal/
│   └── models/
│       ├── errors.go        # Custom error types
│       └── snippets.go      # Snippet model
├── ui/
│   └── html/               # HTML templates
├── .env                    # Environment variables (not in version control)
├── .env.example           # Environment variables template
├── .gitignore             # Git ignore rules
├── go.mod                 # Go module file
├── go.sum                 # Go module checksums
└── README.md              # This file
```

## Security Notes

- Never commit your `.env` file to version control
- Generate secure random strings for `SESSION_SECRET` and `CSRF_SECRET` in production
- Use strong database passwords
- Consider using environment-specific `.env` files (e.g., `.env.production`)

## Development

To contribute to this project:

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment
3. Set up your database according to the instructions above
4. Run the application with `go run ./cmd/web`

## License

This project is for educational purposes.