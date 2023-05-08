package helper

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func BcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(c *gin.Context, email string) (bool, error) {
	var jwtKey = []byte(os.Getenv("SECRET_KEY"))
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return false, err
	}
	AccessExpirationTime := time.Now().Add(1 * time.Minute)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    accessTokenString,
		Expires:  AccessExpirationTime,
		HttpOnly: true,
	})

	return true, nil
}
