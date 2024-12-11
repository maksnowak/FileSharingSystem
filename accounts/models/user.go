package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// # User
//
// A structure representing MongoDB User structure
type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username     string        `bson:"username" json:"username"`
	Email        string        `bson:"email" json:"email"`
	PasswordHash string        `bson:"passwordHash" json:"passwordHash"`
	PasswordSalt string        `bson:"passwordSalt" json:"passwordSalt"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
	Role         string        `bson:"role" json:"role"`
	OwnedFiles   []string      `bson:"ownedFiles" json:"ownedFiles"`
	SharedFiles  []string      `bson:"sharedFiles" json:"sharedFiles"`
}

type Credentials struct {
	Username     string `bson:"username" json:"username"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash"`
}
