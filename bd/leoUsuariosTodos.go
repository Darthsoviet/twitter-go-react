package bd

import (
	"context"
	"log"
	"time"

	"github.com/Darthsoviet/twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoUsuariosTodos lee los usuarios registrados en el sistema, si recibe R trae solo los relacionados con el usuario
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := MongoCN.Database("twittor").Collection("usuarios")

	var res []*models.Usuario

	findOptions := options.Find().SetSkip((page - 1) * 20).SetLimit(20)
	query := bson.M{"nombre": bson.M{"$regex": `(?i)` + search}}

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return res, false, err
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			return res, false, err
		}
		r := models.Relacion{
			UsuarioID:         ID,
			UsuarioRelacionID: s.ID.Hex(),
		}
		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		log.Println(incluir)
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Email = ""

			res = append(res, &s)
		}
	}
	cursor.Close(ctx)
	return res, true, nil
}
