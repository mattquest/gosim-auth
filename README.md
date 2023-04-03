# Authentication Module

This authentication module provides basic JWT authentication and user registration for your Go HTTP server.

## Usage

1. Import the `auth` package in your `main.go` file:

```go
import "github.com/mattquest/gosim-auth/auth"
```

2. Use the provided handlers and middleware in your HTTP server:

```go
package main

import (
	"log"
	"net/http"

	"github.com/mattquest/gosim-auth/auth"
)

func main() {
	http.Handle("/graphql", auth.AuthenticateMiddleware(yourGraphQLHandler))
	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/login", auth.LoginHandler)

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Replace `yourGraphQLHandler` with your actual GraphQL handler or any other handler that requires authentication.
