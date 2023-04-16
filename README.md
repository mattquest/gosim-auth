# Authentication Module

This authentication module provides basic JWT authentication and user registration for your Go HTTP server.

## Optional: Creating a new Go project (Go module)

If you haven't created a Go project yet, follow these steps:

1. Create a new directory for your project:

```sh
mkdir yourproject
cd yourproject
```

Replace `yourproject` with the desired name for your project.

2. Initialize the Go module:

```sh
go mod init github.com/yourusername/yourproject
```

Replace `yourusername` and `yourproject` with your GitHub username and your project's name.

Now you can start using the `auth` package in your Go project.

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
