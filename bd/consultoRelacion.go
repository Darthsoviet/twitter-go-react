package bd

import (
	"context"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion reicbe un modelo de ralacion y verifica si existe en la bd
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := MongoCN.Database("twittor").Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
