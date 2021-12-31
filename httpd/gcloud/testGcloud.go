package gcloud

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Gconnect(c *gin.Context) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error to connection.",
		})
		return
	}
	defer client.Close()
	f, err := os.Open("D:\\learn_golang\\src\\test_api\\httpd\\gcloud\\lisa.jpg")
	if err != nil {
		fmt.Errorf("os.Open: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	defer f.Close()
	ctx, Cancel := context.WithTimeout(ctx, time.Second*30)
	defer Cancel()
	wc := client.Bucket("noseason").Object("lisa").NewWriter(ctx)
	fmt.Println(wc)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Errorf("io.Copy: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if err := wc.Close(); err != nil {
		fmt.Errorf("Writer.Close: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

}

func DownloadImage(c *gin.Context) {
	var img imageWant
	c.BindJSON(&img)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	defer client.Close()
	ctx, Cancel := context.WithTimeout(ctx, 30*time.Second)
	defer Cancel()
	dw, err := client.Bucket("noseason").Object(img.Name).NewReader(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	defer dw.Close()
	data, err := ioutil.ReadAll(dw)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	err = ioutil.WriteFile("lisa.jpg", data, 0)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	// f, err := os.Create("lisa.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// f.Write(data)

}

type imageWant struct {
	Name string `json:"name"`
}

func DeleteImageGcloud(c *gin.Context) {
	var img imageWantDelete
	err := c.BindJSON(&img)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	bucket := client.Bucket("noseason").Object(img.Name)
	if err := bucket.Delete(ctx); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println("delete file " + img.Name + " successful")

}

type imageWantDelete struct {
	Name string `json:"name"`
}
