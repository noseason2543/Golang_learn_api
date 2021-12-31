package handler

import (
	"github.com/gin-gonic/gin"
)

func SumFromDB(c *gin.Context) {
	// client := mongodb.ConnectDB()
	// ctx := context.Background()
	// err := client.Connect(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// idolDB := client.Database("test").Collection("idols")
	// ctx, Cancel := context.WithTimeout(ctx, 40*time.Second)
	// defer Cancel()


	// cursor1, err := idolDB.Distinct(ctx, "name", bson.M{"group": "blackpink"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }


	// cursor2, err := idolDB.CountDocuments(ctx, bson.M{"group": "blackpink"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	
	// cursor3, err := idolDB.Aggregate(ctx, []bson.M{
	// 	{"$match": bson.M{}}, {"$group": bson.M{
	// 		"_id":   "$group",
	// 		"count": bson.M{"$sum": 1},
	// 		"price": bson.M{"$sum": "$price"},
	// 	}},
	// })
	// var s []bson.M

	// if err = cursor3.All(ctx, &s); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// cursor4, err := idolDB.Aggregate(ctx, []bson.M{{"$match": bson.M{}}, {"$lookup": bson.M{
	// 	"from":         "users",
	// 	"localField":   "price",
	// 	"foreignField": "price",
	// 	"as":           "easy",
	// }}, {"$unwind": "$easy"}})

	// var look []bson.M
	// if err = cursor4.All(ctx, &look); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// c.JSON(200, look)
}
