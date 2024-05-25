package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/quilo-bikcodes/Go-JWT/initializers"
	"github.com/quilo-bikcodes/Go-JWT/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//* Get the req body 

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//* Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//* Create User
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)
	
	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//* Respond

	c.JSON(http.StatusBadGateway, gin.H{
		"error": "Failed to create user",
	})
	
}

func Login (c *gin.Context){
	//* Get the req body

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//* Lookup in the DB
	var user models.User
	initializers.DB.First(&user, "Email = ?", body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return		
	}

	//* Compare sent password with saved

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Password",
		})
		return	
		} 

	//* Generate JWT Token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Token",
		})
		return	
	}

	//* Send it back

	c.JSON(http.StatusOK, gin.H{
		"token" : tokenString,
	})
}