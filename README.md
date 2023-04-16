# gosim-auth

A simple authentication package for Go projects, featuring JWT and bcrypt. This package provides user registration, login, and authentication middleware.

## Getting Started

If you haven't created a Go project yet, first create a new directory for your project and initialize it as a Go module by running the following commands:

```sh
mkdir myproject
cd myproject
go mod init github.com/yourusername/myproject
```

Replace `yourusername` and `myproject` with appropriate values.

## Installation

To install the `gosim-auth` package, run the following command in your Go project directory:

```sh
go get github.com/mattquest/gosim-auth
```

## Usage

In your Go project, import the `gosim-auth` package:

```go
import "github.com/mattquest/gosim-auth"
```

### User Registration

To handle user registration, use the `auth.RegisterHandler` function:

```go
http.HandleFunc("/register", auth.RegisterHandler)
```

### User Login

To handle user login, use the `auth.LoginHandler` function:

```go
http.HandleFunc("/login", auth.LoginHandler)
```

### Authentication Middleware

To protect your routes with JWT authentication, use the `auth.AuthenticateMiddleware` function:

```go
http.Handle("/protected", auth.AuthenticateMiddleware(http.HandlerFunc(protectedHandler)))
```

Replace `protectedHandler` with the actual handler function you want to protect.

## License

This project is licensed under the MIT License.
