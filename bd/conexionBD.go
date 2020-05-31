package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el cliente de conexion a la base de datos
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:admin@cluster0-7jbru.mongodb.net/test?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&ssl=true")

/*ConectarBD es la funcion que genera la conexion*/
func ConectarBD() *mongo.Client {

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
