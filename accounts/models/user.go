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
	Username     string        `bson:"username" json:"username" example:"SuzukiSwift"`
	Email        string        `bson:"email" json:"email" example:"suzuki.swift@my.beloved"`
	PasswordHash string        `bson:"passwordHash" json:"passwordHash" example:"f1881511920da79a65164eec5f99b18fd45a4e117b9bcedebe01364899c48d54"`
	PasswordSalt string        `bson:"passwordSalt" json:"passwordSalt" example:"8aa02738f5b472d7c0c4293473ff968cdddd14718be579d8f4ebd792e0f81ef7"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt" example:"2024-12-11T13:58:47.977Z"`
	Role         string        `bson:"role" json:"role" example:"admin"`
	OwnedFiles   []string      `bson:"ownedFiles" json:"ownedFiles" example:"6740a84ebda333c3ff62b0fb,18727a085b5ff06008027b11"`
	SharedFiles  []string      `bson:"sharedFiles" json:"sharedFiles" example:"ef4dc590d2d1356ca4137598,cef1ee14b343dbae4d77020c"`
}

// # Salt
//
// A structure containing the username and password salt associated with it
type Salt struct {
	Username     string `bson:"username" json:"username" example:"ToyotaCorolla"`
	PasswordSalt string `bson:"passwordSalt" json:"passwordSalt" example:"8cf2283ad6ef0a3266059b418a73f8479338233ea2c4bcd3c1f51c39f13ae7dc"`
}

// # Credentials
//
// A structure representing user credentials
type Credentials struct {
	Username     string `bson:"username" json:"username" example:"NissanSkyline"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash" example:"29a52d8ceb75e13713b42509ca3f203372d0cc68bcd797f49538d1ee244d0270"`
}

// # Register
//
// A structure representing data needed to register an account
type Register struct {
	Username     string `bson:"username" json:"username" example:"VolkswagenPassatB5"`
	Email        string `bson:"email" json:"email" example:"1.9TDI@germany.gov"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash" example:"2b2961a431b23c9007efe270c1d7eb79c19d4192d7cd2d924176eb0b19e7d2a1"`
	PasswordSalt string `bson:"passwordSalt" json:"passwordSalt" example:"720f6ac947ab8e448c3a7bb9109e62b72d79a388cb8e05469a6188d76b2e02b3"`
	Role         string `bson:"role" json:"role" example:"admin"`
}

// # Update
//
// A structure representing account data that can be updated
type Update struct {
	Email        string   `bson:"email" json:"email" example:"2.5V6TDI@germany.gov"`
	PasswordHash string   `bson:"passwordHash" json:"passwordHash" example:"15c22c39b6d979a8da5b6185d70a7937630319288cc39fd74a3f705c397281a2"`
	PasswordSalt string   `bson:"passwordSalt" json:"passwordSalt" example:"3246e8fa9472a7c958f3afd81a50acb042ddef9f1558ff3d0ccd4771309aca1f"`
	OwnedFiles   []string `bson:"ownedFiles" json:"ownedFiles" example:"6740a84ebda333c3ff62b0fb,18727a085b5ff06008027b11,bac300aba05c5d3610fcd03b"`
	SharedFiles  []string `bson:"sharedFiles" json:"sharedFiles" example:"ef4dc590d2d1356ca4137598"`
}
