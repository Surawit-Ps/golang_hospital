package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorizes() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		var token string
		if tokenHeader != "" && strings.HasPrefix(tokenHeader, "Bearer ") {
			token = strings.TrimPrefix(tokenHeader, "Bearer ")
		} else {
			cookie, err := c.Cookie("access_token")
			if err == nil {
				token = cookie
			}
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		jwtWrapper := JwtWrapper{
			SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
			Issuer:          "authService",
			ExpirationHours: 24,
		}

		claims, err := jwtWrapper.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// Set the extracted claims in the context
		c.Set("userID", claims.UserID)
		c.Set("HospitalID", claims.HospitalID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func SetCookies(c *gin.Context, token string) {
	c.SetCookie(
		"access_token", // name
		token,          // value
		3600*24,        // maxAge (seconds)
		"/",            // path
		"",             // domain
		false,          // secure
		true,           // httpOnly
	)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
