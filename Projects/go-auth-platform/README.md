# go-auth-platform

### High Level Authentication & Authorization System — Project Documentation

**Built by:** Showrav Kormokar

**Repository:** `github.com/showravkormokar/GO-Language/Projects/go-auth-platform`

---

## 📑 Table of Contents

- [1. Project Overview](#1-project-overview)
- [2. Authentication vs Authorization](#2-authentication-vs-authorization)
- [3. Token and JWT](#3-token-and-jwt)
- [4. Cookie Security](#4-cookie-security)
- [5. Why Access + Refresh Token?](#5-why-access--refresh-token)
- [6. Industry Authentication Architecture](#6-industry-authentication-architecture)
- [7. Architecture Used](#7-architecture-used)
- [8. Permission Matrix](#8-permission-matrix)
- [9. Database Design](#9-database-design)
- [10. API Design](#10-api-design)
- [11. Technologies Used](#11-technologies-used)
- [12. Pagination and Filtering](#12-pagination-and-filtering)
- [13. Performance](#13-performance)
- [14. Bottlenecks](#14-bottlenecks)
- [15. Scaling Plan](#15-scaling-plan)
- [16. Final Achievement](#16-final-achievement)
- [17. Final Summary](#17-final-summary)
- [📘 API Documentation](#-api-documentation)
  - [Health Check](#1-health-check)
  - [Auth Module](#auth-module)
  - [User Module](#user-module)
  - [Admin Module](#admin-module)
  - [Complete Request Flow Architecture](#complete-request-flow-architecture)

---

## 1. Project Overview

`go-auth-platform` is a **production-oriented Authentication and Authorization backend system** built with Go.

The goal was not only to make login/register work, but to build a system similar to what real companies use:

- Secure authentication
- Multi-role authorization
- JWT access/refresh token system
- Token revocation
- Secure cookies
- Role-based access control (RBAC)
- Pagination + filtering
- PostgreSQL optimized queries
- Clean architecture
- Scalable backend design

### What can this system do?

#### Authentication
Users can:
- Register account
- Login
- Logout
- Refresh expired access token
- Change password
- Forgot password flow
- Delete own account

#### Authorization

The system supports:
```
Admin
 ↓
Manager
 ↓
Editor
 ↓
User
```
Different roles have different permissions.

Example:

```
User
 ↓
Only own profile


Manager
 ↓
Manage users


Admin
 ↓
Everything
```

[⬆ back to top](#-table-of-contents)

---

## 2. Authentication vs Authorization

### Authentication

Authentication means: It verifies identity.
> "Who are you?"

### Example:

User enters:
```
email
password
```

System checks:
```
Does this user exist?
Password correct?
```

If yes:
```
Authenticated
```

Real examples:
- Login to Facebook
- Login to Gmail
- Banking login

---

### Authorization

Authorization means:
> "What are you allowed to do?"

### Example:
After login:

User:
```
GET /users/me
```
Allowed.

But:
```
DELETE /admin/users/123
```
Rejected.

Because:
```
User != Admin
```

---

### Why both are needed?

- Without authentication: Anyone can access data.
- Without authorization: Every logged-in user can do everything.

### Example:

A bank:

Authentication:
```
You are Rahim
```

Authorization:
```
Rahim can see his account
but cannot delete another person's account
```

[⬆ back to top](#-table-of-contents)

---

## 3. Token and JWT

### What is Token?

A token is a digital proof.

Instead of sending:
```
email
password
```

Every request: Client sends-
```
Authorization proof
```

### Example:
```
Cookie:
access_token=xxxx
```

Server verifies:
```
Is this token valid?
```

---

### JWT (JSON Web Token)

JWT is a self-contained token format.

Structure:

```
HEADER.PAYLOAD.SIGNATURE
```

### Example:
```
xxxxx.yyyyy.zzzzz
```

#### 1. Header

Contains algorithm:
```json
{
    "alg":"HS256",
    "typ":"JWT"
}
```

#### 2. Payload
Contains claims:
```json
{
    "user_id":"123",
    "role":"admin",
    "exp":1234567
}
```

#### 3. Signature

Created using:
```
secret key
```

### Example:
```
HMACSHA256(
header.payload,
secret
)
```

---

### Why JWT is Stateless?

Traditional session:

Server stores:
```
session_id
 ↓
Database
 ↓
User data
```

JWT:

Server does not store session. Token itself contains information.

Flow:
```
Client
 ↓
JWT
 ↓
Server verifies signature
 ↓
Allow request
```

Advantages:
- No session storage
- Easy horizontal scaling
- Multiple servers can verify same token

[⬆ back to top](#-table-of-contents)

---

## 4. Cookie Security

Cookie stores tokens.

### Example:

```
Set-Cookie:

access_token=abc123
HttpOnly
Secure
SameSite
```

### Why HttpOnly?

JavaScript cannot read it.

Prevents:
```
XSS attack
```

### Example:

Attacker injects:
```javascript
document.cookie
```

Result:
```
access_token unavailable
```

[⬆ back to top](#-table-of-contents)

---

## 5. Why Access + Refresh Token?

A simple JWT:
```
Login
 ↓
Generate JWT
 ↓
Use forever until expiry
```

Problem:
> If stolen: Attacker can use it.

---

### Access Token

Short life:

### Example:
```
15 minutes
```

Used for API access.

### Example:
```
GET /users/me
```

---

### Refresh Token

Long life:

### Example:
```
7 days
```
Used only to create new access token.

Access expired:
```
Access Token ❌

Refresh Token
      ↓
      ↓
New Access Token
```

---

### Token Rotation

Old refresh token is destroyed.

### Example:

Before:
```
Refresh A
```

Request:
```
/refresh
```

After:
```
Refresh A ❌

Refresh B ✅
```

Why?
> If attacker steals old refresh token: It no longer works.

---

### Token Revocation

Means: Immediately disable token.

### Example: User logout-

Before:
```
Access Token
Refresh Token
```

After: Database-
```
blacklisted_tokens

jti
expiry
```

Now:
```
Request
 ↓
JWT
 ↓
Check blacklist
 ↓
Rejected
```

---

### Why single JWT fails?

Problem:
```
JWT expires in 1 hour
```

User logout after 5 minutes. Token still valid for:
```
55 minutes
```

Attacker can use it.
Solution:
- Short access token
- Refresh token
- Rotation
- Revocation

---

### Systems that need this
> Use access/refresh/revocation in:

#### Banking

Example:
- Account data
- Transactions

#### E-commerce

Example:
- Orders
- Payments

#### SaaS applications

Example:
- Multiple users
- Organizations

#### Enterprise systems

Example:
- Admin dashboard
- Employee management

[⬆ back to top](#-table-of-contents)

---

## 6. Industry Authentication Architecture

Architecture:
```
Client
    ↓
    ↓
HTTP Request
    ↓
    ↓
Middleware
    ↓
    ↓
Handler
    ↓
    ↓
Service Layer
    ↓
    ↓
Repository
    ↓
    ↓
PostgreSQL
```

[⬆ back to top](#-table-of-contents)

---

## 7. Architecture Used

### Modular Monolith

Meaning: One application-
```
go-auth-platform
```

but internally separated.
Like:
```
->Auth Module

->User Module

->Role Module

->Token Module
```

Future: Can split easily-
```
->auth-service

->user-service
```

---

### Service + Repository Pattern

#### Handler

Only HTTP work:
- Read request
- Validate
- Return response

#### Service

Business logic:

### Example:
```
check password

generate token

change password
```

#### Repository

Database work:

### Example:
```
Find user

Update user

Create token
```

---

### Folder Structure

```
go-auth-platform/

├── cmd/
│   └── server/
│      └── main.go
├── internal/
|   │
|   ├── config/
|   │   ├── config.go
|   │   ├── env.go
|   │   └── database.go
|   │
|   ├── models/
|   │   ├── user.go
|   │   ├── role.go
|   │   ├── refresh_token.go
|   │   ├── blacklist_token.go
|   │   └── password_reset.go
|   |
|   ├── dto/
|   │   ├── admin/
|   │   ├── auth/
|   │   ├── common/
|   │   ├── paginated/
|   │   └── user/
|   |
|   ├── repository/
|   │   ├── admin_repository.go
|   │   ├── interfaces.go
|   │   ├── user_repository.go
|   │   ├── role_repository.go
|   │   ├── refresh_repository.go
|   │   └── blacklist_repository.go
|   |
|   ├── service/
|   │   ├── auth_service.go
|   │   ├── user_service.go
|   │   ├── errors.go
|   │   ├── password_service.go
|   │   └── admin_service.go
|   |
|   ├── handler/
|   │   ├── auth_handler.go
|   │   ├── user_handler.go
|   │   └── admin_handler.go
|   |
|   ├── middleware/
|   │   ├── auth_middleware.go
|   │   ├── role_middleware.go
|   │   ├── logging.go
|   │   └── recovery.go
|   |
|   ├── migrations/
|   │   ├── migrate.go
|   │   ├── search_index.go
|   │   └── seed.go
|   |
|   ├── routes/
|   │   └── routes.go
|   |
|   ├── utils/
|   │   ├── jwt.go
|   │   ├── password.go
|   │   ├── hash.go
|   │   ├── reset_token.go
|   │   └── response.go
|   |
|   ├── constants/
|   └── tests/
.env
go.mod
Makefile
README.md

```

[⬆ back to top](#-table-of-contents)

---

## 8. Permission Matrix

| Operation          | User | Editor | Manager | Admin |
| ------------------ | ---- | ------ | ------- | ----- |
| View own profile   | ✓    | ✓      | ✓       | ✓     |
| Update own profile | ✓    | ✓      | ✓       | ✓     |
| Change password    | ✓    | ✓      | ✓       | ✓     |
| View all users     | ✗    | ✗      | ✓       | ✓     |
| View user details  | ✗    | ✗      | ✓       | ✓     |
| Deactivate user    | ✗    | ✗      | ✓       | ✓     |
| Assign role        | ✗    | ✗      | ✗       | ✓     |
| Delete user        | ✗    | ✗      | ✗       | ✓     |
| Manage roles       | ✗    | ✗      | ✗       | ✓     |

[⬆ back to top](#-table-of-contents)

---

## 9. Database Design

Database:
```
PostgreSQL
```

>Why relational DB?
Because:
- Users have relationships
- Roles are structured
- Transactions required
- Data consistency important

---

### Tables

#### users

```
id UUID PK
name
email UNIQUE
password_hash
role_id FK
is_active
created_at
updated_at
deleted_at
```

---

#### roles

```
id
name
description
```

---

#### refresh_tokens

```
id
user_id
token_hash
expires_at
revoked
created_at
```

---

#### blacklisted_tokens

```
id
jti UNIQUE
expires_at
created_at
```

---

### Indexing

Index improves search speed.

### Example:

Without index:
```
- 1 million users

- search email

- scan all rows
```

With index:
```
- B-tree index

- direct lookup
```

Used:
```
email
role_id
is_active
created_at
trigram name/email
```

[⬆ back to top](#-table-of-contents)

---

## 10. API Design

REST API:

Means: Using HTTP methods-
```
GET
POST
PATCH
DELETE
```

### Example:

```
GET /users
```
Get users.

---

### API Versioning

We use:
```
/api/v1/
```

Because: Future- 
```
/api/v2/
```
can exist without breaking old clients.

---

### Authentication Flow

```
  Login
    ↓
    ↓
Validate email/password
    ↓
    ↓
bcrypt compare
    ↓
    ↓
Generate:
-> Access Token
        &
-> Refresh Token
    ↓
    ↓
Set HttpOnly Cookie
    ↓
    ↓
Client Request
    ↓
    ↓
Auth Middleware
    ↓
    ↓
Verify JWT
    ↓
    ↓
Check blacklist
    ↓
    ↓
Attach user context
    ↓
    ↓
  Handler
    ↓
    ↓
  Service
    ↓
    ↓
Repository
    ↓
    ↓
Database

```

[⬆ back to top](#-table-of-contents)

---

## 11. Technologies Used

### Go

Why:
- Fast
- Low memory
- Goroutines
- Great for backend

---

### Gorilla/mux

Why:
- URL params
```
/users/{id}
```
- Middleware chaining
- Router groups

---

### GORM
Why:
- ORM
- Relations
- Soft delete
- Transactions
- Preload

---

### bcrypt
Cost:
```
12
```
> Reason: Password cracking becomes expensive.

[⬆ back to top](#-table-of-contents)

---

## 12. Pagination and Filtering

Instead of:
```
GET /users
```

Return 1 million users.

Use:
```
?page=1
&limit=20
```

### Example:

```
page=1

limit=20
```

Database:
```
LIMIT 20
OFFSET 0
```

Benefits:
- Less memory
- Faster response
- Better UX

[⬆ back to top](#-table-of-contents)

---

## 13. Performance

Hardware:
```
4 vCPU
4GB RAM
20GB SSD
```

### Approx:

### Login

bcrypt heavy:
```
50-150 ms
```

Possible:
```
100-300 login/sec
```

---

### Normal API

Example: GET profile-
```
2-10 ms
```

Possible:
```
1000-5000 req/sec
```

---

### Concurrent Users

Go goroutines: Can handle-
```
Thousands of connections
```

Typical:
```
10K+ idle connections
```
depends on database.

[⬆ back to top](#-table-of-contents)

---

## 14. Bottlenecks

### bcrypt

- Intentional slow.
- Security feature.

---

### Blacklist DB check

Each request:
```
SELECT blacklist WHERE jti
```

> Solution: Redis later.

---

### Search

ILIKE: slow.

Solution:
```
pg_trgm
GIN index
```

[⬆ back to top](#-table-of-contents)

---

## 15. Scaling Plan

### Database

Add:
```
Read Replica
```
for GET requests.

---

### Redis
Move blacklist:
```
DB
↓
Redis
```

---

### Multiple Servers

JWT stateless:
```
            User
             ↓
             ↓(Request)
        Load Balancer
             ↓
             ↓
    /--------↓--------\
Server 1  Server 2   Server 3

```

[⬆ back to top](#-table-of-contents)

---

## 16. Final Achievement

After completing this project you learned:

### Backend Engineering

✓ REST API design
✓ Clean architecture
✓ MVC
✓ Repository pattern
✓ Service layer

### Security

✓ JWT
✓ Refresh tokens
✓ Token rotation
✓ Token blacklist
✓ Cookie security
✓ bcrypt

### Database

✓ PostgreSQL
✓ ORM
✓ Relations
✓ Indexing
✓ Pagination

### Production Concepts

✓ RBAC
✓ Middleware
✓ Error handling
✓ Logging
✓ Scaling

[⬆ back to top](#-table-of-contents)

---

## 17. Final Summary

You built:
> A production-style authentication and authorization platform using Go, PostgreSQL, JWT, GORM, and RBAC architecture that can be extended into enterprise SaaS, banking, e-commerce, or microservice-based systems.

This project is no longer a basic login system. It is a foundation-level backend security platform.

[⬆ back to top](#-table-of-contents)

---

---

# 📘 API Documentation

## 1. Health Check

### API Endpoint
```
GET

/health
```
**Access:** Public
**Middleware:** None

### Purpose
Check server status, uptime, memory usage, CPU information.

### Request Body
> No body.

### Response Example

```json
{
    "success": true,
    "message": "Server is healthy and alive",
    "speed_ms": 0,
    "uptime": "2h30m",
    "time": "2026-06-27T10:00:00+06:00",
    "memory_mb": 5,
    "cpu_count": 2,
    "goroutines": 5,
    "version": "1.0.0v"
}
```

### Flow

```
Client
    ↓
Health Handler
    ↓
Collect Runtime Stats
    ↓
Return JSON Response
```

[⬆ back to top](#-table-of-contents)

---

## AUTH MODULE

Base URL:
```
/api/v1/auth
```

---

### 2. Register User

#### API Endpoint
```
POST

/api/v1/auth/register
```
**Access:** Public
**Middleware:** None

#### Purpose
> Create a new user account.

#### Request Body
```json
{
    "name":"Showrav",
    "email":"showrav@gmail.com",
    "password":"12345678"
}
```

#### Flow

```
Handler
    ↓
Decode JSON
    ↓
Validate fields
    ↓
Check existing email
    ↓
Hash password (bcrypt)
    ↓
Create User
    ↓
Assign default role(User)
    ↓
Save PostgreSQL
    ↓
Return response
```

#### Response
```json
{
    "success":true,
    "message":"user created successfully, please login"
}
```

---

### 3. Login

#### API Endpoint
```
POST

/api/v1/auth/login
```
**Access:** Public
**Middleware:** None

#### Purpose
> Authenticate user and generate JWT tokens.

#### Request Body
```json
{
    "email":"showrav@gmail.com",
    "password":"12345678"
}
```

#### Flow

```
Handler
    ↓
Validate fields
    ↓
Service
    ↓
Find user by email
    ↓
Check active status
    ↓
bcrypt password verify
    ↓
Generate Access Token
    ↓
Generate Refresh Token
    ↓
Store refresh token hash
    ↓
Set HttpOnly Cookie
    ↓
Return User Data
```

#### Response

```json
{
    "success":true,
    "message":"login success",
    "data":{
        "id":"uuid",
        "name":"Showrav",
        "email":"showrav@gmail.com",
        "role":"user"
    }
}
```

Cookies:
```
access_token = JWT
refresh_token = JWT
```

---

### 4. Refresh Token

#### API Endpoint
```
POST

/api/v1/auth/refresh
```
**Access:** Public
**Middleware:** None

#### Purpose
> Generate new access token using refresh token.

#### Request Body
> No body. Refresh token comes from cookie.

#### Flow

```
Handler
    ↓
Read refresh_token cookie
    ↓
Hash token
    ↓
Find token DB
    ↓
Check expiry
    ↓
Check revoked status
    ↓
Generate new access token
    ↓
Rotate refresh token
    ↓
Update cookie
    ↓
Response
```

#### Response
```json
{
    "success":true,
    "message":"token refreshed successfully"
}
```

---

### 5. Forgot Password

#### API Endpoint
```
POST

/api/v1/auth/forgot-password
```
**Access:** Public
**Middleware:** None

#### Purpose
> Create password reset request.

#### Request Body
```json
{
    "email":"showrav@gmail.com"
}
```

#### Flow

```
Handler
    ↓
Validate email
    ↓
Find user
    ↓
Generate random reset token
    ↓
Hash token
    ↓
Store DB
    ↓
Create reset link
    ↓
Send email (or return dummy link for development)
    ↓
Response
```

#### Response
```json
{
    "success":true,
    "message":"If email exists, reset link sent"
}
```

---

### 6. Reset Password

#### API Endpoint
```
POST

/api/v1/auth/reset-password
```
**Access:** Public
**Middleware:** None

#### Request Body
```json
{
    "token":"a03009d0983...",
    "new_password":"newPassword123"
}
```

#### Flow

```
Handler
    ↓
Validate token
    ↓
Hash token
    ↓
Find reset token
    ↓
Check expiry
    ↓
Check used status
    ↓
Hash new password
    ↓
Update user password
    ↓
Mark token used
    ↓
Revoke all sessions
    ↓
Return response
```

#### Response
```json
{
    "success":true,
    "message":"password reset successfully"
}
```

[⬆ back to top](#-table-of-contents)

---

## USER MODULE

Base:
```
/api/v1/users
```

Protected:
```
AuthRequired Middleware
```

---

### 7. Logout

#### API Endpoint
```
POST

/api/v1/logout
```
**Access:** Private
**Middleware**
```
AuthRequired
```

#### Purpose
> Logout current device.

#### Request
> No body.

#### Flow

```
Middleware
    ↓
Verify JWT
    ↓
Handler
    ↓
Service
    ↓
Revoke Refresh Token
    ↓
Blacklist Access Token
    ↓
Clear Cookies
    ↓
Response
```

#### Response

```json
{
    "success":true,
    "message":"logout success"
}
```

---

### 8. Get Current User Profile

#### API Endpoint
```
GET

/api/v1/users/me
```

**Access:** Private
**Middleware:**
```
AuthRequired
```

#### Request
> No body

#### Flow

```
Middleware
    ↓
Decode JWT
    ↓
Attach user id context
    ↓
Handler
    ↓
Service
    ↓
Find user
    ↓
Preload role
    ↓
Return User
```

#### Response

```json
{
    "success":true,
    "data":{
        "id":"uuid",
        "name":"Showrav",
        "email":"showrav@gmail.com",
        "role":{
            "name":"user"
        }
    }
}
```

---

### 9. Change Password

#### API Endpoint
```
PATCH

/api/v1/users/me/password
```

**Access:** Private
**Middleware:**
```
AuthRequired
```

#### Request Body
```json
{
    "current_password":"old123456",
    "new_password":"new123456"
}
```

#### Flow

```
Handler
    ↓
Validate fields
    ↓
Get current user id
    ↓
Service
    ↓
Find user
    ↓
bcrypt verify old password
    ↓
Hash new password
    ↓
Update password
    ↓
Revoke all refresh tokens
    ↓
Blacklist current token
    ↓
Response
```

#### Response

```json
{
    "success":true,
    "message":"password changed successfully, please login"
}
```

[⬆ back to top](#-table-of-contents)

---

## ADMIN MODULE

Base:
```
/api/v1/admin
```

Middleware:
```
AuthRequired
+
RequireMinRole(admin)
```

---

### 10. Get All Users

#### API Endpoint
```
GET

/api/v1/admin/users
```

Query:
```
?page=1
&limit=20
&search=john
&role=editor
&is_active=true
&sort=created_at
&order=desc
```

#### Flow

```
Middleware
    ↓
Admin Permission Check
    ↓
Handler
    ↓
Parse filters
    ↓
Service
    ↓
Repository
    ↓
GORM Query
    ↓
Pagination
    ↓
Return Data
```

#### Response
```json
{
    "success":true,
    "data":[
        {
            "name":"John",
            "email":"john@gmail.com"
        }
    ],
    "pagination":{
        "page":1,
        "limit":20
    }
}
```

---

### 11. Get Specific User

#### API Endpoint
```
GET

/api/v1/admin/users/{id}
```

### Example:
```
/api/v1/admin/users/uuid
```

#### Flow

```
Admin Middleware
    ↓
Handler
    ↓
Validate UUID
    ↓
Service
    ↓
Repository FindByID
    ↓
Preload Role
    ↓
Return User
```

#### Response

```json
{
    "success":true,
    "data":{
        "id":"uuid",
        "name":"John",
        "role":{
            "name":"editor"
        }
    }
}
```

---

### 12. Assign User Role

#### API Endpoint
```
PATCH

/api/v1/admin/users/{id}/role
```

#### Body
```json
{
    "role_id":373
}
```

#### Flow

```
Admin Middleware
    ↓
Validate Role
    ↓
Find User
    ↓
Update RoleID
    ↓
Save
    ↓
Return User
```

#### Response

```json
{
    "success":true,
    "message":"role updated successfully"
}
```

---

### 13. Update User Status

#### API Endpoint
```
PATCH

/api/v1/admin/users/{id}/status
```

Permission:
```
Manager + Admin
```

Middleware:
```
AuthRequired
+
RequireMinRole(manager)
```

#### Body
```json
{
    "is_active":false
}
```

#### Flow

```
Middleware
    ↓
Check Permission
    ↓
Handler
    ↓
Service
    ↓
Update IsActive
    ↓
Save
    ↓
Return User
```

#### Response

```json
{
    "success":true,
    "message":"user status updated"
}
```

---

### 14. Update User

#### API Endpoint

```
PATCH

/api/v1/admin/users/{id}
```

Permission:
```
Admin
```

#### Body
```json
{
    "name":"New Name",
    "email":"new@gmail.com",
    "role_id":282,
    "is_active":true
}
```

#### Flow

```
Admin Middleware
    ↓
Validate fields
    ↓
Create update map
    ↓
Update only changed fields
    ↓
Revoke sessions
    ↓
Save
    ↓
Response
```

#### Response

```json
{
    "success":true,
    "message":"user updated successfully"
}
```

---

### 15. Delete User

#### API Endpoint
```
DELETE

/api/v1/admin/users/{id}
```

Permission:
```
Admin
```

#### Flow

```
Admin Middleware
    ↓
Find User
    ↓
Soft Delete
    ↓
Revoke refresh tokens
    ↓
Blacklist sessions
    ↓
Return 204
```

Response:
```
HTTP 204 No Content
```

---

### 16. Get All Roles

#### API Endpoint

```
GET

/api/v1/admin/roles
```

Permission:
```
Admin
```

#### Flow

```
Middleware
    ↓
Handler
    ↓
Service
    ↓
Role Repository
    ↓
Return Roles
```

#### Response

```json
{
    "success":true,
    "data":[
        {
            "id":191,
            "name":"admin"
        },
        {
            "id":464,
            "name":"user"
        }
    ]
}
```

[⬆ back to top](#-table-of-contents)

---

## Complete Request Flow Architecture

For Protected API:

```
Client
    ↓
Route
    ↓
Auth Middleware
    ↓
JWT Verify
    ↓
Blacklist Check
    ↓
Role Middleware
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
GORM
    ↓
PostgreSQL
    ↓
Response
```

This is the same flow used in production-grade Go backend systems.

[⬆ back to top](#-table-of-contents)
