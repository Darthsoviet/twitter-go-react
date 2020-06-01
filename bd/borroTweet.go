package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BorroTweet borra un tweet de la base de datos
func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	col := MongoCN.Database("twittor").Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err

}
