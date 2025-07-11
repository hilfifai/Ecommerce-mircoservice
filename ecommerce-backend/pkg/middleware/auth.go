package middleware

import (
	"ecommerce-backend/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/golang-jwt/jwt/v5"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func AuthMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip middleware untuk preflight requests dan endpoint publik
		if c.Request.Method == "OPTIONS" || isPublicEndpoint(c.Request.URL.Path) {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			return
		}

		// Validasi format token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token format",
				"message": "Format harus: Bearer <token>",
			})
			return
		}

		tokenString := tokenParts[1]
		
		// Validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token",
				"message": "Token tidak valid atau sudah kadaluarsa",
			})
			return
		}

		// Ekstrak claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token claims",
				"message": "Gagal memproses token",
			})
			return
		}

		// Validasi issuer dan audience jika diperlukan
		if !claims.VerifyIssuer("your-app-name", true) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token issuer",
				"message": "Issuer token tidak valid",
			})
			return
		}

		// Set context dengan data user
		c.Set("userID", claims["sub"])
		c.Set("userRole", claims["role"])
		c.Set("userData", claims)
		
		c.Next()
	}
}

func isPublicEndpoint(path string) bool {
	publicEndpoints := []string{
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/products",
	}
	
	for _, endpoint := range publicEndpoints {
		if strings.HasPrefix(path, endpoint) {
			return true
		}
	}
	return false
}