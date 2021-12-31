package handler

import (
	"context"
	"log"
	"net/http"
	mongodb "test_apigg/httpd/mongoDB"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginTest(c *gin.Context) {
	var u loginUser
	var resultFromFind FindInDB
	var check bool
	c.BindJSON(&u)
	client := mongodb.ConnectDB()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	UserDB := client.Database("test").Collection("users")

	if err := UserDB.FindOne(ctx, bson.M{"email": u.Email}).Decode(&resultFromFind); err != nil {
		log.Fatal(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(resultFromFind.PasswordDB), []byte(u.Password)); err != nil {

		check = false
	} else {
		check = true
	}

	if check == true {
		claims := &Claims{
			ID:       resultFromFind.ID,
			Email:    resultFromFind.EmailDB,
			Password: resultFromFind.PasswordDB,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"token": t,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "password is wrong",
		})
	}

}

type Claims struct {
	ID       string
	Email    string
	Password string
	jwt.StandardClaims
}

type FindInDB struct {
	ID         string `bson:"_id"`
	EmailDB    string `bson:"email"`
	PasswordDB string `bson:"password"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
