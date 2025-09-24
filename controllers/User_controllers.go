package controllers

import (
	"time"

	"github.com/ayushwar/ecommerce/database"
	"github.com/ayushwar/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func Register(ctx *gin.Context)  {
	var user models.User
	if err:=ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(400,gin.H{"error":"invaild request"})
		return
	}
	var existinguser models.User

	if err:=database.DB.Where("email=?",user.Email).First(&existinguser).Error;err==nil{
		ctx.JSON(400,gin.H{"error":"user already exist"})
		return
	}
	hashed,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost);
	if err!=nil{
		ctx.JSON(400,gin.H{"error":"invaild rewuest"})
		return
	}
	user.Password=string(hashed)
	if err:=database.DB.Create(&user).Error;err!=nil{
		ctx.JSON(404,gin.H{"error":"hashed error"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User registered successfully"})
	
}
// JWT claims structure
type JWTClaims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

func Login(ctx *gin.Context)  {
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err:=ctx.ShouldBindJSON(&input);err!=nil{
		ctx.JSON(400,gin.H{"err":"invaild request"})
		return
	}
	var user models.User
	if err:=database.DB.Where("email=?",input.Email).First(&user).Error;err!=nil{
		ctx.JSON(400,gin.H{"error":"user not found"})
		return
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password));err!=nil{
		ctx.JSON(400,gin.H{"invalid":"password"})
		return
	}
	claims:=JWTClaims{
		UserID: user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "ayush",

		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtSecret)
	if err!=nil{
		ctx.JSON(400,gin.H{"Error":"failed to gen token"})
		return
	}
	   ctx.JSON(200, gin.H{
        "message": "Login successful",
        "token":   tokenString,
    })
}