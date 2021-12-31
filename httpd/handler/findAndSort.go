package handler

import (
	"context"
	"fmt"
	mongodb "test_apigg/httpd/mongoDB"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAndSort(c *gin.Context) {
	ctx := context.Background()
	client := mongodb.ConnectDB()
	if err := client.Connect(ctx); err != nil {
		fmt.Println(err)
		return
	}
	find := client.Database("test").Collection("idols")
	ctx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	allIdols, err := find.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{"group", 1}, {"name", 1}}))
	if err != nil {
		fmt.Println(err)
		return
	}
	var result []bson.M
	if err := allIdols.All(ctx, &result); err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, result)
}
