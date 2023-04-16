# gosim-auth

A simple Go authentication package for user registration, login, and JWT authentication.

## Prerequisites

1. Install Go: https://golang.org/doc/install
2. Make sure you have a new or existing Go project and it is a Go module.

## How to use gosim-auth

1. Install the package:

```
   go get github.com/mattquest/gosim-auth
```

2. Import the package in your project:

```
   import "github.com/mattquest/gosim-auth"
```

3. Define a User Store that implements the auth.UserStore interface:

```
   type YourUserStore struct {
   }

   func (s *YourUserStore) GetUserByEmail(email string) (*auth.User, error) {
       // Your implementation here
   }

   func (s *YourUserStore) CreateUser(user *auth.User) error {
       // Your implementation here
   }
```

4. Initialize the User Store and handlers:

```
   userStore := &YourUserStore{}
   registerHandler := auth.NewRegisterHandler(userStore)
   loginHandler := auth.NewLoginHandler(userStore)
   authenticateMiddleware := auth.AuthenticateMiddleware
```

5. Use the handlers in your HTTP server:

```
   http.HandleFunc("/register", registerHandler)
   http.HandleFunc("/login", loginHandler)
   http.Handle("/protected", authenticateMiddleware(protectedHandler))
```
