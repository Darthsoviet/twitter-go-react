package bd

import (
	"context"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificoRegistro actualiza usuario
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	colection := MongoCN.Database("twittor").Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apllidos) > 0 {
		registro["apellidos"] = u.Apllidos
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Email) > 0 {
		registro["email"] = u.Email
	}
	if len(u.Password) > 0 {
		registro["password"] = u.Password
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID}}

	_, err := colection.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
