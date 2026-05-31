# JWT Authentication + User Profile API (MongoDB) with Go

This document is the learning roadmap, setup guide, and design reference for building a **User Management API** with **JWT authentication**, **MongoDB**, and **Go (Golang)**. Completing this project will give you core backend skills: **Middleware**, **Authentication**, **Authorization**, **Password Hashing**, and **Protected Routes**.

---

## Why this project?

- It covers **real-world backend concepts** in one cohesive project.
- You'll build a complete **authentication flow** (register → login → JWT → protected routes).
- You'll learn **MongoDB integration**, **environment configuration**, and **clean project structure**.
- It's a solid foundation for any production API (e.g., e‑commerce, SaaS, social apps).

---

## Goals

By the end of Phase‑1, you will be able to:

- Implement **Register**, **Login**, **JWT generation**, and **JWT verification**.
- Build **Get Profile**, **Update Profile**, and **Delete Profile** endpoints.
- Enforce **protected routes** with **JWT middleware**.
- Implement **role-based authorization** (`Admin` vs `User`).
- Store users securely with **bcrypt password hashing**.
- Connect Go to **MongoDB** and perform CRUD operations.
- Use **environment variables** for configuration.
- Follow a **clean, scalable project structure**.

---

## Outcomes

After completing this project, you will have:

- A working **HTTP server** in Go with **JWT authentication**.
- A **MongoDB-backed user store**.
- A fully functional **auth API** with **protected routes**.
- Experience with **middleware**, **authentication flow**, and **authorization**.
- A reusable template for future Go backend projects.

---

## Concepts: Shortly about Go Backend

Go (Golang) is a compiled, statically-typed language known for:

- **High performance** and efficiency (close to C in many cases).
- **Simple syntax** and fast compilation.
- **Built-in concurrency** (goroutines, channels).
- A rich standard library for **HTTP servers**, **JSON**, and **I/O**.
- Strong tooling: `go mod`, `go build`, `go test`, `go fmt`.

For web APIs, Go is often used with:

- **gorilla/mux** for routing.
- **MongoDB Go driver** for database access.
- **JWT libraries** for authentication.
- **bcrypt** for password hashing.
- **godotenv** for environment variable management.

---

## Learning Roadmap (Next Step)

### Phase‑1 (Current Target)

**Build: User Management API**

- [ ] **Register**
- [ ] **Login**
- [ ] **JWT Generate**
- [ ] **JWT Verify Middleware**
- [ ] **Get Profile**
- [ ] **Update Profile**
- [ ] **Delete Profile**
- [ ] **Protected Routes**
- [ ] **Role Authorization (Admin/User)**

---

## Best Practice Authentication Flow

```text
# Phase01
Register
   ↓
Password Hash (bcrypt)
   ↓
Store User

# Phase02
Login
   ↓
Compare Password
   ↓
Generate JWT Access Token

# Phase03
Client
   ↓
Store Token

# Phase04
Protected API
   ↓
Authorization: Bearer <token>

# Phase05
Middleware
   ↓
Verify JWT
   ↓
User Context
   ↓
Handler
```

### For Learning (Simpler)

- Store token in **`localStorage`** (easy for learning and prototyping).

### For Real Production (Recommended)

- **Access Token** → **HttpOnly Cookie**
- **Refresh Token** → **HttpOnly Cookie**

HttpOnly cookies are more secure against XSS attacks.

---

## Project Structure

```text
httpServer_MongoDB_JWT/

│ main.go
├── config/
│   └── env.go
├── database/
│   └── mongo.go
├── models/
│   └── user.go
├── dto/
│   ├── authRequest.go
│   └── authResponse.go
├── services/
│   ├── authService.go
│   └── userService.go
├── middleware/
│   └── jwtMiddleware.go
├── utils/
│   ├── jwt.go
│   └── password.go
├── routes/
│   └── routes.go
├── types/
│   └── response.go
└── .env
```

This structure separates:

- **Configuration** (`config/`)
- **Database logic** (`database/`)
- **Data models** (`models/`)
- **Request/Response shapes** (`dto/`, `types/`)
- **Business logic** (`services/`)
- **Middleware** (`middleware/`)
- **Utilities** (`utils/`)
- **Routing** (`routes/`)
- **Entry point** (`main.go`)

---

## Install Dependencies

### 1. Create Project

```bash
mkdir httpServer_MongoDB_JWT
cd httpServer_MongoDB_JWT
go mod init httpServer_MongoDB_JWT
```

### 2. Install Packages

```bash
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
```

These provide:

- **gorilla/mux**: HTTP router and middleware.
- **MongoDB Go driver**: MongoDB connectivity.
- **golang-jwt/jwt/v5**: JWT creation and validation.
- **bcrypt**: Secure password hashing.
- **godotenv**: Load `.env` files.

---

## Environment Variables

Create a `.env` file in the project root:

```env
PORT=8000
MONGO_URI=mongodb://localhost:27017
DB_NAME=userStore
JWT_SECRET=my_super_secret_key
```

---

## API Endpoints Summary

| Method | Endpoint        | Description               | Auth Required |
|--------|-----------------|---------------------------|---------------|
| POST   | `/register`     | Register new user         | No            |
| POST   | `/login`        | Login & get JWT           | No            |
| GET    | `/api/profile`  | Get user profile          | Yes           |
| PUT    | `/api/profile`  | Update user profile       | Yes           |
| DELETE | `/api/profile`  | Delete user profile       | Yes           |
| GET    | `/api/admin/users` | List users (admin only) | Yes (Admin)   |

Authorization header format:

```http
Authorization: Bearer <JWT_TOKEN>
```

---

## Next Steps After Phase‑1

- Add **refresh token** flow with HttpOnly cookies.
- Add **email verification** and **password reset**.
- Add **input validation** (e.g., using `go-playground/validator`).
- Add **unit tests** and **integration tests**.
- Add **logging**, **error handling middleware**, and **structured logs**.
- Containerize with **Docker** and deploy.

---

## Summary

This project teaches you:

- **JWT Authentication** (generate, verify, use in protected routes).
- **Password Hashing** with bcrypt.
- **MongoDB integration** in Go.
- **Middleware** for authentication and authorization.
- **Role-based access control** (Admin vs User).
- A **clean, production-like project structure**.
- How to use **environment variables** and **configuration**.

Completing this will give you a strong foundation in **Go backend development** and **secure API design**.