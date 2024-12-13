package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	FileID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FileName  string             `json:"file_name,omitempty" bson:"fileName,omitempty"`
	UserID    string             `json:"user_id,omitempty" bson:"userID,omitempty"`
	Tags      []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Data      []byte             `json:"data,omitempty" bson:"data,omitempty"`
	HasAccess []string           `json:"has_access,omitempty" bson:"hasAccess,omitempty"` // List of user IDs
}
