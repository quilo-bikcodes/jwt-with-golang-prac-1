package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/quilo-bikcodes/Go-JWT/initializers"
	"github.com/quilo-bikcodes/Go-JWT/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("\033[38;5;208m","In middleware...","\033[0m")   //Extra Code for Orange Color 

	//* Get the cookie off req body

	tokenString, err := c.Cookie("Authorization")

	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//* Decode-Validate It

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	//* Check the exp

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix())  > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
			}	
		

	

	//* Find the user with token subject

	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//* Attach Req
	c.Set("user", user)

	//* Continue

	c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}