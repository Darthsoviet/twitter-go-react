package bd

import (
	"context"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoTweet  graba tweet en bd
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := MongoCN.Database("twittor").Collection("tweet")

	// registro := bson.M{
	// 	"userID":  t.UserID,
	// 	"mensaje": t.Mensaje,
	// 	"fecha":   t.Fecha,
	// }
	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
