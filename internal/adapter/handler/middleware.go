package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/util"
)

// AuthMiddleware validates JWT token from cookie or Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt_token")
		if err != nil {
			// fallback: check Authorization header
			authHeader := c.GetHeader("Authorization")
			tokenString, _ = strings.CutPrefix(authHeader, "Bearer ")
			// if strings.HasPrefix(authHeader, "Bearer ") {
			// 	tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			// }
		}

		if tokenString == "" {
			util.ResponseHandler(c, http.StatusUnauthorized, true, "Unauthorized")
			return
		}

		// parse token
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			util.ResponseHandler(c, http.StatusUnauthorized, true, "Invalid or expired token")
			return
		}

		// extract claims
		claims, ok := token.Claims.(*Claims)
		if !ok {
			util.ResponseHandler(c, http.StatusUnauthorized, true, "Invalid token claims")
			return
		}

		// store user claims in context for downstream handlers
		c.Set("user", claims)
		c.Next()
	}
}

// ensures that the user has the required role
func RoleMiddleware(requiredRole domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			util.ResponseHandler(c, http.StatusForbidden, true, "No user in context")
			return
		}

		claims := user.(*Claims)
		if claims.Role != requiredRole {
			util.ResponseHandler(c, http.StatusForbidden, true, "Forbidden")
			return
		}

		c.Next()
	}
}
