package models

//Tweet Captrua el body de la peticion, el mensaje que nos llega
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
