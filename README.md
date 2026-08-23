# Go is an open source programming language that makes it simple to build secure, scalable systems. #
# Simple Server Using Go

A backend web server and RSS feed aggregator built with Go as part of Go programming practice. This project includes user management, RSS feeds, feed following, and a built-in background scraper.

## Features

- **RESTful Endpoints:** Handle user creation, authentication, feeds, and feed follows.
- **Database Integration:** Uses SQL for data persistence with generated type-safe code via `sqlc`.
- **Authentication Middleware:** Secure specific routes using API keys/tokens.
- **Background Scraper:** Automatically fetches and processes RSS feeds concurrently.

## Project Structure

```text
.
├── internal/           # Internal application packages
├── sql/                # SQL migrations and queries
├── vendor/             # Vendor dependencies
├── handler_feed.go         # RSS feed handlers
├── handler_feed_follows.go # Feed follow relationship handlers
├── handler_readiness.go    # Health check endpoint handler
├── handler_user.go         # User management handlers
├── json.go                 # JSON helper functions
├── main.go                 # Application entry point
├── middleware_auth.go      # Authentication middleware
├── models.go               # Data models
├── scraper.go              # Background RSS feed scraper
└── sqlc.yaml               # SQLC configuration file
```


```text
Refer: `vendor` directory for modules
```
## Credits: FreeCodeCamp.org
 Video: https://youtu.be/un6ZyFkqFKo?si=v-jzJB2WKbemXj-J 
