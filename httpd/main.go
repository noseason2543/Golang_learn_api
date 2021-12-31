package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test_apigg/httpd/gcloud"
	"test_apigg/httpd/handler"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://noseason:password@cluster0.4ejho.mongodb.net/test?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dbFind := client.Database("test").Collection("users")
	// result, err := dbFind.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer result.Close(ctx)

	// var re []bson.M
	// if err = result.All(ctx, &re); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(re)
	d := []int{1, 2, 3}
	// var j int = len(d)
	//เอาไปใช้กับ กำหนด array ไม่ได้เพราะ มันเป็นค่าไม่แน่นอน
	var q [2]*int
	q[0] = &d[0]
	q[1] = &d[1]

	for i, v := range q {
		fmt.Println(*q[i])
		fmt.Println(*v)
	}

	r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	// r.GET("/album", GetAlbums)
	// r.GET("/", func(c *gin.Context) {
	// 	c.IndentedJSON(http.StatusOK, re)
	// })
	r.GET("/sign", HomeQuery)
	r.GET("/sign/:name/:age", HomeParams)
	r.POST("/", PostHomePage)
	r.POST("/signUp", handler.SignUpNewUser)
	r.POST("/login", handler.LoginTest)
	r.POST("/getUser", handler.GetDetailUser)
	r.GET("/testGoogle", gcloud.Gconnect)
	r.POST("/testDownload", gcloud.DownloadImage)
	r.POST("/testDelete", gcloud.DeleteImageGcloud)
	r.POST("/search", handler.Search)
	r.GET("/countD", handler.SumFromDB)
	r.GET("/findAndSort", handler.FindAndSort)
	r.Run(":3000")

	
}

func HomeQuery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func HomeParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func PostHomePage(c *gin.Context) {
	var u User
	// c.BindJSON(&u)  // ใช้ BindJSON จะง่ายกว่ามานั่ง unmarshal
	input := c.Request.Body
	body, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &u)
	fmt.Println(u.Name + " " + u.Age)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://noseason:password@cluster0.4ejho.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dbFind := client.Database("test").Collection("users")
	findResult, err := dbFind.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var re []bson.M
	if err = findResult.All(ctx, &re); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, re)
}

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{{Id: "noseason", Title: "P'Tu", Artist: "hereTu", Price: 3000.0},
	{Id: "booker", Title: "devin", Artist: "deva", Price: 3200.0}}

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
