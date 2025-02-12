package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/agrawal-2005/iiita_lab/pkg/models"
)

// Secret key for JWT
var JWT_SECRET = []byte(os.Getenv("SECRET_KEY"))

// AuthController struct
type AuthController struct {
	DB *gorm.DB
}

// Register function
func (ac *AuthController) Register(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Type       string `json:"type" binding:"required"`
		Email      string `json:"email" binding:"required"`
		Name       string `json:"name" binding:"required"`
		Department string `json:"department"`
		Domain     string `json:"domain,omitempty"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Store Login credentials
	login := models.Login{
		Username: input.Username,
		Password: string(hashedPassword),
		Type:     input.Type,
	}
	if err := ac.DB.Create(&login).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create login credentials"})
		return
	}

	// Store People details
	people := models.People{
		Email:      input.Email,
		Name:       input.Name,
		Department: input.Department,
		Type:       input.Type,
		Domain:     input.Domain,
	}
	if err := ac.DB.Create(&people).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login function
func (ac *AuthController) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.Login

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	if err := ac.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": "Login successful",
	})
}

// GetProfile function
func (ac *AuthController) GetProfile(c *gin.Context) {
	email := c.Param("email")

	var profile models.People
	if err := ac.DB.Where("email = ?", email).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// Logout function
func (ac *AuthController) Logout(c *gin.Context) {
	// Invalidate the token by instructing the client to remove it
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
		"note":    "Please remove the token from client storage.",
	})
}
