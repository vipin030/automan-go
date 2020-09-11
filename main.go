package main

import (
	"github.com/vipin030/automan/controllers"
	"github.com/vipin030/automan/models"
	"os"
	"net/http"
	"fmt"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.POST("/login", controllers.Authenticate)
	r.GET("/vtypes", TokenAuthMiddleware(),controllers.FindVehicleTypes)
	r.GET("/vtypes/:id", controllers.FindVehicleType)
	r.POST("/vtypes", controllers.CreateVehicleType)
	r.PATCH("/vtypes/:id", controllers.UpdateVehicleType)
	r.DELETE("/vtypes/:id", controllers.DeleteVehicleType)

	// Run the server
	r.Run(":8080")
}

func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  //normally Authorization the_token_xxx
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
  tokenString := ExtractToken(r)
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     //Make sure that the token method conform to "SigningMethodHMAC"
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
     return []byte(os.Getenv("ACCESS_SECRET")), nil
  })
  if err != nil {
     return nil, err
  }
  return token, nil
}

func TokenValid(r *http.Request) error {
  token, err := VerifyToken(r)
  if err != nil {
     return err
  }
  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
     return err
  }
  return nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
     err := TokenValid(c.Request)
     if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        c.Abort()
        return
     }
     c.Next()
  }
}
