

---

# 🛒 Golang E-Commerce Backend

**A simple, scalable E-commerce backend built with Golang, Gin, GORM, and MySQL.**

---

## 🚀 Project Overview

This project is a backend system for an E-commerce platform.
It allows users to:

* Register and login (JWT authentication)
* Browse products
* Add products to cart
* Place orders
* Admins can manage products

The backend is built with **Golang** and uses **MySQL** for database storage.

---

## 🧰 Tech Stack

* **Language:** Golang (Go)
* **Web Framework:** Gin
* **Database ORM:** GORM
* **Database:** MySQL
* **Authentication:** JWT
* **Password Security:** bcrypt

---

## 📦 Features

### User Management

* User registration and login
* JWT-based authentication
* Role-based access (`user` or `admin`)

### Product Management

* Create, update, delete products (admin)
* Get all products (public)

### Cart

* Add products to cart
* View cart items
* Remove items from cart

### Orders

* Place order from cart
* View user orders
* Order total calculated automatically
* Order items stored with snapshot price

---

## 🔑 Environment Variables

Create a `.env` file in the project root:

```env
DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=ecommerce_db
JWT_SECRET=your-secret-key
```

> **Important:** Do not commit `.env` to GitHub (add it to `.gitignore`).

---

## ⚙️ Installation

1. Clone the repository:

```bash
git clone https://github.com/<your-username>/ecommerce-golang.git
cd ecommerce-golang
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create MySQL database:

```sql
CREATE DATABASE ecommerce_db;
```

4. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

---

## 📚 API Endpoints

### Public Routes

| Method | Endpoint  | Description             |
| ------ | --------- | ----------------------- |
| POST   | /register | Register a new user     |
| POST   | /login    | User login, returns JWT |

### Protected Routes (JWT required)

| Method | Endpoint               | Description            |
| ------ | ---------------------- | ---------------------- |
| GET    | /api/products          | Get all products       |
| POST   | /api/products          | Create product (admin) |
| PUT    | /api/products/\:id     | Update product (admin) |
| DELETE | /api/products/\:id     | Delete product (admin) |
| POST   | /api/cart              | Add item to cart       |
| GET    | /api/cart/\:user\_id   | Get user cart          |
| DELETE | /api/cart/\:id         | Remove item from cart  |
| POST   | /api/orders            | Place order from cart  |
| GET    | /api/orders/\:user\_id | Get all user orders    |

> Use the `Authorization: Bearer <token>` header for protected routes.

---

## 🔐 Authentication

* Passwords are hashed using `bcrypt`.
* JWT is used for authentication.
* Tokens expire after 24 hours.

---

## 📂 Project Structure

```
ecommerce/
├─ controllers/   # Route handlers for users, products, cart, orders
├─ middlewares/   # JWT auth middleware
├─ models/        # GORM models (User, Product, Order, Cart)
├─ database/      # Database connection and migration
├─ routes/        # Gin routes
├─ main.go        # Entry point
├─ .env           # Environment variables (not committed)
```

---

## 💡 Notes

* Prices are stored in **paise/cents** to avoid float precision issues.
* Cart items and order items are related to products using foreign keys.
* Orders snapshot product prices at the time of purchase.

---

## 📌 Future Improvements

* Payment gateway integration (Stripe/PayPal)
* Product categories and filters
* Order status updates (shipped, delivered, cancelled)
* Pagination for product listings
* Admin dashboard for managing users and orders

---

## ⚡ Author

**Ayush Ahirwar (AICE)**
Backend project built in Golang for resume and portfolio.

---


Do you want me to add that?
