package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DevuelvoTweetSeguidores Modelo de weets seguidores
type DevuelvoTweetSeguidores struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioRelacion string             `bson:"usuariorelacionid," json:"userRelationId,omitempty"`
	UsuarioID       string             `bson:"usuarioid" json:"usuarioId,omitempty"`
	Tweet           struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
