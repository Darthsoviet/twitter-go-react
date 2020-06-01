package bd

import (
	"context"
	"log"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweet devuleve un a lista de tweets
func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	col := MongoCN.Database("twittor").Collection("tweet")
	defer cancel()

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{"userid": ID}
	opciones := options.Find().SetLimit(20).SetSort(bson.D{{Key: "fecha", Value: -1}}).SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)

	}
	return resultados, true
}
