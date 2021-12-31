package handler

import (
	"context"
	"fmt"
	mongodb "test_apigg/httpd/mongoDB"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Search(c *gin.Context) {
	var input inputFromUser
	err := c.BindJSON(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	clientDB := mongodb.ConnectDB()
	ctx := context.Background()
	err = clientDB.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	memberDB := clientDB.Database("test").Collection("idols")
	ctx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	cursorIdols, err := memberDB.Find(ctx, bson.M{"name": primitive.Regex{Pattern: input.Input, Options: "i"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	brandDB := clientDB.Database("test").Collection("brands")
	cursorBrand, err := brandDB.Find(ctx, bson.M{"brand": primitive.Regex{Pattern: input.Input, Options: "i"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	groupDB := clientDB.Database("test").Collection("groups")
	cursorGroup, err := groupDB.Find(ctx, bson.M{"group": primitive.Regex{Pattern: input.Input, Options: "i"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	var resultIdols []bson.M
	var resultBrand []bson.M
	var resultGroup []bson.M
	if err := cursorBrand.All(ctx, &resultBrand); err != nil {
		fmt.Println(err)
		return
	}
	if err = cursorIdols.All(ctx, &resultIdols); err != nil {
		fmt.Println(err)
		return
	}
	if err = cursorGroup.All(ctx, &resultGroup); err != nil {
		fmt.Println(err)
		return
	}

	result := append(append(resultBrand, resultGroup...), resultIdols...)
	c.JSON(200, result)

}

type inputFromUser struct {
	Input string `json:"input"`
}
