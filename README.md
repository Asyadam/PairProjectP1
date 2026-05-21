# 🎮 Game Store Management CLI

CLI (Command Line Interface) application using Golang and MySQL for managing a physical game store.

This system is designed to help store admins manage:
- User authentication
- Game inventory
- Categories
- Transactions / Orders
- Reports


---

# 🚀 Main Features

## 🔐 Authentication System
Users must login before accessing the system.

Features:
- Register
- Login
- Logout

---

## 🎮 Game Management
Manage game inventory in the store.

Features:
- Add Game
- View Games
- Update Game
- Delete Game

---

## 🏷️ Category Management
Manage game categories.

Features:
- Add Category
- View Categories
- Update Category
- Delete Category

Examples:
- Action
- RPG
- Horror
- Sports

---

## 🛒 Order / Transaction System
Manage customer purchases.

Features:
- Create Order
- Add Games to Order
- Calculate Total Price
- View Transactions

---

## 📊 Reports
Generate reports for store management.

Reports:
- User Reports
- Sales Reports
- Revenue Reports
- Stock Reports
- Most Purchased Games

---

# 👥 Team Members

-Aidil Syadam
-Silvanus

---

# 🔥 Task Distribution

## 👨‍💻 Aidil Syadam

Responsible for:
- Authentication System
- Main Menu
- Game CRUD
- Order / Transaction System
- Database Connection
- Main Application Flow

Files:

```bash
handler/auth_handler.go
repository/auth_repository.go

handler/game_handler.go
repository/game_repository.go

handler/order_handler.go
repository/order_repository.go

cli/menu.go
cli/game_menu.go
cli/order_menu.go

cmd/main.go
db/db.go
```

---

## 👨‍💻 Silvanus

Responsible for:
- Category CRUD
- Reports Feature
- Unit Testing
- Documentation

Files:

```bash
handler/category_handler.go
repository/category_repository.go

handler/report_handler.go
repository/report_repository.go

cli/category_menu.go
cli/report_menu.go

tests/
README.md
docs/
```

---

# 📂 Project Structure

```bash
PairProjectP1/
│
├── cmd/
├── db/
├── entity/
├── repository/
├── handler/
├── cli/
├── tests/
├── docs/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

# 🌳 Branch Workflow

## Main Branch

```bash
main
```

Stable production branch.

---

## Feature Branches

```bash
feature/core-system
feature/order-system
feature/report-category
```

---

# 🗄️ Database Explanation

Database Name:

```sql
gamestore
```

The database is designed using:
- One To One Relationship
- One To Many Relationship
- Many To Many Relationship

---

# 📌 TABLE EXPLANATION

---

## 👤 users

Stores login account information.

Columns:
- id
- email
- password
- role
- created_at

Relationship:
- One user has one profile
- One user can have many orders

---

## 🪪 profiles

Stores additional user information.

Columns:
- id
- user_id
- full_name
- phone
- address

Relationship:
- One To One with users

---

## 🎮 games

Stores game inventory.

Columns:
- id
- title
- price
- stock
- description
- release_date

Relationship:
- One game can belong to many categories
- One game can appear in many order details

---

## 🏷️ categories

Stores game categories.

Columns:
- id
- category_name

Examples:
- RPG
- Horror
- Adventure
- Sports

---

## 🔗 game_categories

Bridge table for Many To Many relationship.

Purpose:
- One game can have many categories
- One category can contain many games

Example:
- GTA V → Action
- GTA V → Adventure

---

## 🧾 orders

Stores transaction information.

Columns:
- id
- user_id
- order_date
- total_price
- status

Relationship:
- One user can have many orders

---

## 📦 order_details

Stores purchased game details inside an order.

Columns:
- id
- order_id
- game_id
- quantity
- subtotal

Purpose:
- Stores list of purchased games
- Calculates subtotal for each game

Example:
Order #1:
- FIFA 25 x2
- GTA V x1

---

# 🔗 Database Relationships

## ✅ One To One

```text
users → profiles
```

One user has one profile.

---

## ✅ One To Many

```text
users → orders
```

One user can have many orders.

---

## ✅ Many To Many

```text
games ↔ categories
```

Handled using:
```text
game_categories
```

---

# ⚙️ Setup Project

## Install Dependencies

```bash
go mod tidy
```

---

## Setup Environment

Create `.env`

```env
MYSQL_DSN=root:password@tcp(localhost:xxxx)/gamestore
```

---

# ▶️ Run Application

```bash
go run cmd/main.go
```

---

