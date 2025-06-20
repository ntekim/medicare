# Receptionist & Doctor Portal - Golang Web Application

## 📌 Overview

A web application built in Go for managing patient records. It includes role-based portals for **Receptionists** and **Doctors**, with a shared login API and secure access control.

---

## 🧩 Features

- ✅ **Single Login API** for both Receptionists and Doctors.
- 🩺 **Doctor Portal**:
  - View registered patient details.
  - Update patient records (excluding delete).
- 💼 **Receptionist Portal**:
  - Register new patients.
  - Full CRUD on patient records.
- 🔐 **Authentication** with role-based access control (JWT).
- 🧪 **Unit tests** for major logic.
- 🧾 **Swagger Documentation**.
- 🚀 Creative enhancements (optional: e.g. consultation records, audit logs, vitals tracking).

---

## 🛠️ Tech Stack

| Layer        | Technology     |
|--------------|----------------|
| Backend      | Golang (Gin / Chi / Fiber / GoFrame) |
| Database     | PostgreSQL     |
| ORM / Queries| SQLC or GORM (flexible) |
| Auth         | JWT            |
| API Docs     | Swagger / Postman |
| Optional UI  | React / HTML-CSS |
| Migrations   | `golang-migrate`, `sql-migrate`, or `goose` |

---


---

## 🔐 Authentication

- Uses JWT for authentication.
- Role-based access using middleware:
  - `"receptionist"` role required for patient registration & CRUD.
  - `"doctor"` role required for viewing/updating patient data.

---

## 📃 API Endpoints

### Auth

| Method | Endpoint       | Description      |
|--------|----------------|------------------|
| POST   | `/login`       | Login for both doctors & receptionists |

### Patients

| Method | Endpoint        | Description              | Role         |
|--------|-----------------|--------------------------|--------------|
| POST   | `/patients`     | Create patient           | Receptionist |
| GET    | `/patients`     | List all patients        | Receptionist/Doctor |
| GET    | `/patients/{id}`| View a patient           | Both         |
| PUT    | `/patients/{id}`| Update patient           | Both         |
| DELETE | `/patients/{id}`| Delete patient           | Receptionist |

> Full Swagger documentation available at `/swagger/index.html` (if Swagger is enabled)

---

## 🧾 API Documentation

This project uses [OpenAPI](https://www.openapis.org/) and Swagger UI for API documentation.

### 📘 View Swagger UI

- [Local OpenAPI Doc](http://localhost:8000/api.json)
- [Local Swagger Docs](http://localhost:8000/swagger/)



## 🧪 Testing

- Basic unit tests for:
  - Auth service
  - Patient CRUD logic
- Use `go test ./...` to run all tests.

---

## 🚀 How to Run

```bash
# 1. Setup DB (PostgreSQL)

# 2. Run migrations
make migration-up


# 3. Build and run the app
go run main.go

# 4. Generate Sqlc files
make sqlc-gen
