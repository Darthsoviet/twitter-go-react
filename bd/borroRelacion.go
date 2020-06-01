package bd

import (
	"context"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
)

//BorroRelacion borra un documento de la colecion relacion
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := MongoCN.Database("twittor").Collection("relacion")

	_, err := collection.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil

}
