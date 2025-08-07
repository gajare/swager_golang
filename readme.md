**Student Microservice with Swagger** in Golang, aimed at **both end users and developers**.

---

## 📘 `README.md` – Student CRUD Microservice in Golang with Swagger

---

# 📚 Student CRUD Microservice (Golang + PostgreSQL + Swagger)

A simple microservice built in **Go** using **Gorilla Mux**, **PostgreSQL**, and documented with **Swagger UI**.
It allows users to perform CRUD operations on student data and provides a full OpenAPI/Swagger specification for easy understanding and testing.

---

## 📌 Features

* 🧑‍🎓 Student Model: `id`, `name`, `email`, `age`
* ⚙️ RESTful APIs: Create, Read, Update, Delete
* 🗃️ PostgreSQL as the database
* 📑 Auto-generated Swagger UI for API documentation
* ✨ Clean folder structure (models, controller, service, router, db)

---

## 🚀 For End Users: How to Run This Project

### ⚙️ Prerequisites

* Go 1.18+ installed
* PostgreSQL installed or Docker
* `swag` CLI tool installed:

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

---

### 🛠️ Step-by-Step Setup

#### 1. Clone the repository

```bash
git clone https://github.com/gajare/swager_golang.git
cd swager_golang
```

#### 2. Start PostgreSQL

**Option A: Local PostgreSQL**

```sql
CREATE DATABASE books;
```

**Option B: Docker (recommended)**

```bash
docker run --name pg-books -e POSTGRES_PASSWORD=Amol -e POSTGRES_DB=books -p 5432:5432 -d postgres
```

#### 3. Set up database schema manually (if needed)

```sql
CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  name TEXT,
  email TEXT,
  age INT
);
```

#### 4. Insert sample data (optional)

```sql
INSERT INTO students (name, email, age) VALUES
('Aarav Mehta', 'aarav.mehta@example.com', 20),
('Saanvi Sharma', 'saanvi.sharma@example.com', 21),
...
```

#### 5. Generate Swagger docs

```bash
swag init
```

This will create a `/docs` folder with Swagger annotations.

---

#### 6. Run the project

```bash
go run main.go
```

Server will start at: `http://localhost:8080`

---

### ✅ API Endpoints

| Method | Endpoint         | Description          |
| ------ | ---------------- | -------------------- |
| GET    | `/students`      | Get all students     |
| GET    | `/students/{id}` | Get a student by ID  |
| POST   | `/students`      | Create a new student |
| PUT    | `/students/{id}` | Update a student     |
| DELETE | `/students/{id}` | Delete a student     |

---

### 📖 View Swagger UI

📌 Visit:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Use **Try it out** buttons to test live APIs.

---

## 👨‍💻 For Developers: Learning Swagger in a Go Microservice

### 📌 What You'll Learn

If you're a Go developer building microservices and want to integrate Swagger:

* ✅ How to document routes using swaggo annotations
* ✅ How to generate OpenAPI JSON/YAML from Go comments
* ✅ How to serve Swagger UI alongside your API
* ✅ How to test and debug APIs visually using Swagger

---

### 🧱 Key Components

| File/Folder         | Purpose                                      |
| ------------------- | -------------------------------------------- |
| `controller/`       | Contains API handlers with Swagger comments  |
| `service/`          | Business logic layer                         |
| `models/`           | GORM model definition (`Student`)            |
| `db/db.go`          | PostgreSQL connection setup                  |
| `routers/router.go` | Route registration + Swagger UI mount        |
| `docs/`             | Auto-generated Swagger specs via `swag init` |
| `main.go`           | Entry point + server start                   |

---

### 🧩 Example Swagger Comment

```go
// @Summary Get all students
// @Description Fetches all student records
// @Tags students
// @Produce json
// @Success 200 {array} models.Student
// @Router /students [get]
```

> These comments are parsed by `swag init` to generate OpenAPI docs.

---

### ✅ Good to Know

* Swagger JSON is available at `/swagger/doc.json`
* No manual Swagger YAML/JSON writing — everything is auto-generated
* Easy for frontend or QA teams to consume and test API

---

## 🙋‍♂️ Need Help?

Feel free to reach out or open an issue if:

* Swagger UI doesn’t load
* Database connection fails
* You want Docker support (can be added)

---

## 📜 License

MIT License

---