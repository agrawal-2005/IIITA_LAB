# **IIITA Lab**  

## 📌 Project Overview  
IIITA Lab is built using **Golang (Gin Framework)** and **GORM** for database management. It manages **research-related data**, including publications, projects, conferences, students, and supervisors at **IIIT Allahabad**.

## 🚀 Features  
✔️ **User Authentication** (JWT-based login system)  
✔️ **CRUD Operations** for research publications, projects, and conferences  
✔️ **Student and Supervisor Management**  
✔️ **Feedback Collection**  
✔️ **Automatic Database Migration** with GORM  
✔️ **Secure Password Hashing** using bcrypt  
✔️ **Role-Based Access Control (RBAC)**  

---

## 🛠️ Tech Stack  
**Backend:** Golang (Gin Framework)  
**Database:** MySQL (with GORM ORM)  
**Authentication:** JWT (JSON Web Token)  
**Environment Management:** `godotenv`  
**Dependency Management:** `go mod`  

---

## 📂 Project Structure  
```
iiita_lab/
│── backend/                     # Backend folder
│   ├── cmd/                      # Entry point for the app
│   │   ├── main.go               # Main application file
│   ├── pkg/                      # Package folder for configurations, models, and utilities
│   │   ├── config/               # Configuration files (DB, JWT, etc.)
│   │   │   ├── db.go             # MySQL database connection
│   │   ├── controllers/          # API Controllers
│   │   │   ├── student.go        # Student controller
│   │   │   ├── project.go        # Project controller
│   │   ├── middlewares/          # Middleware (Auth, Logging, etc.)
│   │   │   ├── auth.go           # JWT Authentication Middleware
│   │   ├── models/               # Structs representing database tables
│   │   │   ├── models.go         # Database models
│   │   ├── routes/               # API Routes
│   │   │   ├── routes.go         # Route definitions
│   │   ├── utils/                # Utility functions
│   │   │   ├── jwt.go            # JWT handling
│   │   │   ├── password.go       # Password hashing utilities
│   ├── go.mod                    # Go module file
│   ├── go.sum                    # Dependency lock file
│── document/                     # Documentation folder
│   ├── ER.pdf                    # Entity Relationship Diagram
│── frontend/                      # Frontend folder
│   ├── views/                     # Frontend views (HTML, EJS, etc.)
│── README.md                      # Project documentation
│── .env                           # Environment configuration file
```

---

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
JWT_SECRET="your_jwt_secret"
```

### 4️⃣ Install Dependencies  
```sh
go mod tidy
```

### 5️⃣ Run the Application  
```sh
go run backend/cmd/main.go
```

---

## 🔥 API Endpoints  

### **🛡️ Authentication**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| POST   | `/login` | User login (JWT authentication) |

### **📚 Students & Supervisors**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| GET    | `/students` | Retrieve all students  |
| GET    | `/supervisors` | Retrieve all supervisors |

### **📄 Projects**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| GET    | `/projects` | Fetch all projects  |
| POST   | `/projects` | Create a new project |

### **📑 Publications**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| GET    | `/publications` | Fetch all research publications  |

### **📅 Conferences**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| GET    | `/conferences` | Retrieve all conferences  |

### **📝 Feedback**
| Method | Endpoint  | Description  |
|--------|----------|--------------|
| POST   | `/feedback` | Submit user feedback  |

---

## ✨ Contributing  
Contributions are **welcome**! Follow these steps to contribute:  
1. **Fork** the repository  
2. **Create a new branch** (`git checkout -b feature-branch`)  
3. **Make your changes** and commit (`git commit -m "Add new feature"`)  
4. **Push to branch** (`git push origin feature-branch`)  
5. **Open a Pull Request** 🚀  

---

## 📜 License  
This project is licensed under the [MIT License](LICENSE.txt).

🚀 Developed & Maintained by Prashant Agrawal 🚀
Let me know if you need any further modifications! 😊