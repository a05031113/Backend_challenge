package controllers

import (
	"Rehasaku_code_challenge/backend/database"
	"Rehasaku_code_challenge/backend/helper"
	"Rehasaku_code_challenge/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()

func Register(c *gin.Context) {
	var register models.Register
	if err := c.BindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong format"})
		return
	}

	if register.Email == "" || register.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "miss data"})
		return
	}

	var check int64
	db.Table("users").Where("email = ?", register.Email).Count(&check)
	if check > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exist"})
		return
	}

	bcryptPass, _ := helper.BcryptPassword(register.Password)
	register.Password = bcryptPass

	db.Table("users").Select("email", "password").Create(&register)
	c.JSON(http.StatusOK, gin.H{"register": true})
}

func Login(c *gin.Context) {
	var login models.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong format"})
		return
	}

	var check models.Login
	db.Table("users").Where("email = ?", login.Email).Take(&check)
	if check.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
		return
	}

	if match := helper.CheckPasswordHash(login.Password, check.Password); !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
		return
	}

	JWTStatus, err := helper.GenerateJWT(c, login.Email)
	if !JWTStatus {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"login": true})
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, gin.H{"logout": true})
}

func User(c *gin.Context) {
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{"email": email})
}
