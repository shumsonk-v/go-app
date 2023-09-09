package controllers

import (
	"fmt"
	"go-app/inits"
	"go-app/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	// Define form object should look like
	var body struct {
		Username    string `form:"username" json:"username" binding:"required"`
		Email       string `form:"email" json:"email" binding:"required,email"`
		Password    string `form:"password" json:"password" binding:"required"`
		DisplayName string `form:"display_name" json:"display_name" binding:"required"`
		Firstname   string `form:"firstname" json:"firstname" binding:"required"`
		Middlename  string `form:"middlename" json:"middlename"`
		Lastname    string `form:"lastname" json:"lastname" binding:"required"`
	}

	// If requested body failed to bind with form 'body' above, error will be returned
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Begin transaction
	tx := inits.DB.Begin()

	// Instantiate new user from model with requested data
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash), DisplayName: body.DisplayName, Firstname: body.Firstname, Middlename: body.Middlename, Lastname: body.Lastname, Role: models.Normal, Status: models.Pending}

	// Use tx (transaction) to insert new user data
	result := tx.Create(&user)

	if result.Error != nil {
		tx.Rollback()
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	tx.Commit()
	ctx.JSON(200, gin.H{"data": user})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           user.ID,
		"displayName":  user.DisplayName,
		"role":         user.Role,
		"profileImage": user.ProfileImage,
		"exp":          time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	ctx.JSON(200, gin.H{
		"token": tokenString,
	})
}

func GetUsers(ctx *gin.Context) {
	var users []models.User

	err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting users"})
		return
	}

	ctx.JSON(200, gin.H{"data": users})
}

func GetAuthenticatedUser(ctx *gin.Context) {
	user, err := ctx.Get("user")

	if err != false {
		fmt.Println(user)
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"data": "You are logged in!", "user": user})
}
