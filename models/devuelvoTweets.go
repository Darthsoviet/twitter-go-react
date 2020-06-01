package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DevuelvoTweets modelo de tweets
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje"`
	Fecha   time.Time          `bson:"fecha" json:"fecha"`
}
