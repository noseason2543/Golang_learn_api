package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func jwtDetailUser(token, id string) (*jwt.Token, error) {
	tokee, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(id), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("jwt is expired")
	}
	return tokee, nil

}

func GetDetailUser(c *gin.Context) {
	var token reciveToken
	c.BindJSON(&token)
	result, err := jwtDetailUser(token.Token, "secret")
	if err == nil {
		c.JSON(http.StatusOK, result.Claims)
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

type reciveToken struct {
	Token string `json:"token"`
}
