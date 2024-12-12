package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// # User
//
// A structure representing MongoDB User structure
type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id" example:"675f9a97ca1d148373316ae4"`
	Username     string        `bson:"username" json:"username" example:"Karol_Wojtyla"`
	Email        string        `bson:"email" json:"email" example:"huan.pablo.dos@vatican.city"`
	PasswordHash string        `bson:"passwordHash" json:"passwordHash" example:"Kremowki"`
	PasswordSalt string        `bson:"passwordSalt" json:"passwordSalt" example:"Slony_Karmel"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt" example:"2024-12-11T13:58:47.977Z"`
	Role         string        `bson:"role" json:"role" example:"admin"`
	OwnedFiles   []string      `bson:"ownedFiles" json:"ownedFiles" example:"rower,pies"`
	SharedFiles  []string      `bson:"sharedFiles" json:"sharedFiles" example:"zaba,slon"`
}

type Credentials struct {
	Username     string `bson:"username" json:"username" example:"Jon_Bon_Jovi"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash" example:"ZbazowaneDane123"`
}
