# simple-server

A basic starting project for how I learned to setup Go HTTP servers using clean architecture patterns. 

## Libraries Used

- [**chi**](github.com/go-chi/chi/v5) - Simplest HTTP router for quality of life without adding heavy features
- [**slog**](https://pkg.go.dev/log/slog) - Go's standard structured logging package.
- [**godotenv**](github.com/joho/godotenv) - Loads environment variables from .env files. 

## Project Architecture

The project uses clean architecture patterns to separate functionality and maintain clear boundaries between layers.

### Layer Descriptions

#### cmd/
The application's main entry point. 

#### internal/config/
- Loads environment variables from .env files
- Validates required configuration
- Initializes the application logger

#### internal/server/
- Dependencies are created/initialized 
- The HTTP router (chi) is configured
- Middleware is applied 
- Routes are registered and mapped to handlers
- Server lifecycle is managed (start/stop)

#### internal/handler/
- Receives HTTP requests and parses input
- Calls appropriate service methods

#### internal/service/
- Implements core application functionality
- Coordinates between handlers and repositories
- Seperates HTTP concerns 

#### internal/domain/
- Contains pure business logic for specific domains
- Keeps services small by extracting domain operations

#### internal/repository/
- Abstracts data storage operations 
- Implements interfaces defined in the service layer


## API Endpoints

- `GET /v1/` - API version information
- `GET /v1/health` - Health check endpoint
- `GET /v1/test/test` - Basic handler test
- `GET /v1/test/testService` - Service layer test
- `GET /v1/test/testDomain` - Domain layer test
- `GET /v1/test/testRepo` - Repository layer test
