package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Require(c *gin.Context) {
	accessTokenString, err := c.Cookie("accessToken")
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("email", claims["email"])
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
