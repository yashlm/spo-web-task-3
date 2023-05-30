package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   string `gorm:"unique"`
	Password string
}

type SignupRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=postgresyash dbname=task3 port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the User model
	db.AutoMigrate(&User{})

	// Create a new Gin router
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Define routes
	router.POST("/signup", signupHandler)
	router.POST("/login", loginHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(router.Run(":8080"))
}

func signupHandler(c *gin.Context) {
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user already exists
	var existingUser User
	if result := db.Where("user_id = ?", request.UserID).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new user
	newUser := User{UserID: request.UserID, Password: string(hashedPassword)}
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := SignupResponse{Message: "Signup successful"}
	c.JSON(http.StatusOK, response)
}

func loginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user
	var user User
	result := db.Where("user_id = ?", request.UserID).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	response := LoginResponse{Message: "Login successful"}
	c.JSON(http.StatusOK, response)
}
