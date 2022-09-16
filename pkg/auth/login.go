package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pradiptadh/majoo/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequestBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func AuthLogin(c *gin.Context) {
	var req RegisterRequestBody
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	err = db.Where("user_name = ? ", req.UserName).First(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get secret token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	//create cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"cookie": "cookies added"})

}
