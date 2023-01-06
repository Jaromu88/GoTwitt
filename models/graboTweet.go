package models

import "time"

/*GuardarTweet es el formato o estructura que tendrá nuestro tweet en la BD */
type GuardarTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
