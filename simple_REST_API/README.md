# 📝 Simple REST API in Go


A beginner-friendly RESTful API built using Golang, supporting CRUD operations for Users and Posts. Data is stored in memory (maps), making it lightweight and easy to understand.

## 🚀 Features

- Manage Users : Create, Read (all & by ID), Update, Delete
- Manage Posts : Create, Read (all & by ID), Update, Delete
- JSON-based requests and responses
- Delete tasks by ID
- In-memory storage (no database required)


## 🛠️ Built With

- Go – programming language
- net/http – HTTP server & routing
- encoding/json – JSON encoding & decoding
- In-memory store (map) – lightweight persistence
## Installation


1. Clone the Repository

```bash
git clone https://github.com/Shashivarunreddy/Go_projects.git
cd Go_projects/simple_REST_API
```

2. Install dependencies:

```bash
go mod tidy
```
3. Build and install to your Go bin: 


```bash
go run .

```



## 🎯 Example Workflow

1.Users
```bash
# Create Users
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'

curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Bob","email":"bob@example.com"}'

# Get All Users
curl http://localhost:8080/users

# Get User by ID
curl http://localhost:8080/users/1

# Update User
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice@new.com"}'

# Delete User
curl -X DELETE http://localhost:8080/users/1


```

2.Posts
```bash
# Create Posts
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Post","content":"Hello World!"}'

# Get All Posts
curl http://localhost:8080/posts

# Get Post by ID
curl http://localhost:8080/posts/1

# Update Post
curl -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","content":"Updated Content"}'

# Delete Post
curl -X DELETE http://localhost:8080/posts/1

```
