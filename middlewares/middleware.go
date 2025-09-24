package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWTs
var secretKey = []byte("your-secret-key")

// VerifyToken validates the JWT token string
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

// AuthMiddleware protects routes using JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "authorization header missing"})
			ctx.Abort()
			return
		}

		// Format should be "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(401, gin.H{"error": "invalid authorization header format"})
			ctx.Abort()
			return
		}

		// Verify token
		claims, err := VerifyToken(tokenParts[1])
		if err != nil {
			ctx.JSON(401, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		// Attach claims to context (so handlers can use them)
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("role", claims["role"])

		ctx.Next() // continue to next handler
	}
}
