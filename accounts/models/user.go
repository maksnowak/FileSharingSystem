package models

import "time"

// # User
//
// A structure representing MongoDB User structure
type User struct {
	ID           string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string    `bson:"username" json:"username"`
	Email        string    `bson:"email" json:"email"`
	PasswordHash string    `bson:"passwordHash" json:"passwordHash"`
	PasswordSalt string    `bson:"passwordSalt" json:"passwordSalt"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
	Role         string    `bson:"role" json:"role"`
	OwnedFiles   []string  `bson:"ownedFiles" json:"ownedFiles"`
	SharedFiles  []string  `bson:"sharedFiles" json:"sharedFiles"`
}
