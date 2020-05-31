package bd

import (
	"context"
	"log"

	"github.com/Darthsoviet/twitter-go-react/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el cliente de conexion a la base de datos
var MongoCN = ConectarBD()

/*ConectarBD es la funcion que genera la conexion*/
func ConectarBD() *mongo.Client {

	var clientOptions = options.Client().ApplyURI(util.ObtenerEnv("MONGO"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion exitosa con la bd")
	return client
}

/*ChequeoConnection se encarga de confirmar si hay una conexion a la base de datos */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0

	}
	return 1

}
