package bd

import (
	"context"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
)

// InsertoRelacion inserta una relacion en la base de datos
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	colection := MongoCN.Database("twittor").Collection("relacion")

	_, err := colection.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
