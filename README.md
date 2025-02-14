# IIITA Lab API

## 📌 Project Overview
IIITA Lab API is a RESTful API built using **Golang (Gin Framework)** and **GORM** for database management. It provides an interface for managing research-related data, including publications, projects, conferences, students, and supervisors at IIITA.

## 🚀 Features
- User Authentication (Login system)
- CRUD operations for research publications, projects, and conferences
- Student and Supervisor management
- Feedback collection
- Automatic database migration with GORM

## 🛠️ Tech Stack
- **Backend:** Golang (Gin Framework)
- **Database:** MySQL (with GORM ORM)
- **Environment Management:** `godotenv`
- **Dependency Management:** `go mod`

## 📂 Project Structure
```
iiita_lab/
│── backend/                     # Backend folder
│   ├── cmd/                      # Entry point for app
│   │   ├── main.go               # Main application file
│   ├── pkg/                      # Package folder for configurations, models, and utilities
│   │   ├── config/               # Configuration files (DB, JWT, etc.)
│   │   │   ├── db.go             # MySQL database connection
│   │   ├── controllers/          # Handles API logic
│   │   │   ├── student.go        # Student controller
│   │   ├── middlewares/         # Middleware (Auth, Logging, etc.)
│   │   │   ├── auth.go           # JWT Authentication Middleware
│   │   ├── models/               # Structs representing database tables
│   │   │   ├── models.go         # Database models
│   │   ├── routes/               # API Routes
│   │   ├── utils/                # Utility functions
│   │   │   ├── jwt.go            # JWT handling
│   │   │   ├── password.go       # Password hashing utilities
│   ├── go.mod                    # Go module file
│   ├── go.sum                    # Dependency lock file
│── document/                     # Documentation folder
│   ├── ER.pdf                    # Entity Relationship Diagram
│── frontend/                     # Frontend folder
│   ├── views/                     # Frontend views (HTML, EJS, etc.)
│── README.md                      # Project documentation
```

## ⚙️ Setup Instructions
### 1️⃣ Prerequisites
- Install **Go** (version 1.18+ recommended)
- Install **MySQL** and create a database

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/agrawal-2005/iiita_lab.git
cd iiita_lab
```

### 3️⃣ Set Up Environment Variables
Create a `.env` file in the root directory and add:
```sh
DB_USER="your_username"
DB_PASSWORD="your_password"
DB_HOST="localhost"
DB_PORT="3306"
DB_NAME="iiita_lab_db"
```

### 4️⃣ Install Dependencies
```sh
go mod tidy
```

### 5️⃣ Run the Application
```sh
go run backend/cmd/main.go
```

### 6️⃣ API Endpoints
| Method | Endpoint        | Description                     |
|--------|---------------|---------------------------------|
| GET    | `/`           | Welcome message                 |
| POST   | `/login`      | User authentication             |
| GET    | `/students`   | Retrieve student details        |
| POST   | `/projects`   | Create a new project entry      |
| GET    | `/publications` | Fetch all research publications |

## ✨ Contributing
Contributions are welcome! To contribute:
1. Fork the repo
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit changes (`git commit -m "Add feature"`)
4. Push to branch (`git push origin feature-branch`)
5. Open a pull request

## 📜 License  
This project is licensed under the [MIT License](LICENSE.txt).

---
🚀 **Developed & Maintained by Prashant Agrawal** 🚀