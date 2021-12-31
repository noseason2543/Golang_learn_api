package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	mongodb "test_apigg/httpd/mongoDB"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUpNewUser(c *gin.Context) {
	var u newUser
	c.BindJSON(&u)
	fmt.Println(u.Email + " " + u.Password)
	client := mongodb.ConnectDB()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test").Collection("users")
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	password := string(hash)
	insertUser, err := db.InsertOne(ctx, bson.M{"email": u.Email, "password": password})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

type newUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
