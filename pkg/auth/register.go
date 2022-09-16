package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pradiptadh/majoo/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequestBody struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func AuthRegister(c *gin.Context) {

	var req RegisterRequestBody
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	tx := db.Begin()
	defer tx.Rollback()

	newUser := &models.User{
		Name:     req.Name,
		UserName: req.UserName,
		Password: string(hash),
	}
	if err := tx.Create(&newUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"acknowledge": "success create user"})

}
