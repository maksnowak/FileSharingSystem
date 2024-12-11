package models

import "time"

type User struct {
	ID           string    `bson:"_id,omitempty" json:"id,omitempty"`
	username     string    `bson:"username" json:"username"`
	email        string    `bson:"email" json:"email"`
	passwordHash string    `bson:"passwordHash" json:"-"`
	passwordSalt string    `bson:"passwordSalt" json:"-"`
	createdAt    time.Time `bson:"createdAt" json:"createdAt"`
	role         string    `bson:"role" json:"role"`
	ownedFiles   []string  `bson:"ownedFiles" json:"ownedFiles"`
	sharedFiles  []string  `bson:"sharedFiles" json:"sharedFiles"`
}
