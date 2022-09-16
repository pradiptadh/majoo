package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pradiptadh/majoo/pkg/models"
	"gorm.io/gorm"
)

func RequireAuth(c *gin.Context) {

	//get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		return
	}

	//decode or validate

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		}

		var user models.User
		db := c.MustGet("db").(*gorm.DB)
		err = db.First(&user, claims["sub"]).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.Set("user", user)

		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})

	}

}
