package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	FileID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FileName  string             `json:"file_name,omitempty" bson:"fileName,omitempty"`
	UserID    string             `json:"user_id,omitempty" bson:"userID,omitempty"`
	Tags      []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Data      []byte             `json:"data,omitempty" bson:"data,omitempty"`
	HasAccess []string           `json:"has_access,omitempty" bson:"hasAccess,omitempty"` // List of user IDs
}

func (f *File) createFile(ctx *context.Context, collection *mongo.Collection) error {
	_, err := collection.InsertOne(*ctx, f)
	if err != nil {
		return err
	}

	return nil
}

func getAllFiles(ctx *context.Context, collection *mongo.Collection) ([]File, error) {
	var files []File
	cursor, err := collection.Find(*ctx, nil)
	if err != nil {
		return files, err
	}

	err = cursor.All(*ctx, &files)
	if err != nil {
		return files, err
	}

	return files, nil
}

func (f *File) getFile(ctx *context.Context, collection *mongo.Collection) error {
	err := collection.FindOne(*ctx, f).Decode(f)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) updateFile(ctx *context.Context, collection *mongo.Collection) error {
	filter := bson.M{"_id": f.FileID}

	update := bson.M{
		"$set": bson.M{
			"fileName":  f.FileName,
			"userID":    f.UserID,
			"tags":      f.Tags,
			"data":      f.Data,
			"hasAccess": f.HasAccess,
		},
	}

	err := collection.FindOneAndUpdate(*ctx, filter, update).Decode(f)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) deleteFile(ctx *context.Context, collection *mongo.Collection) error {
	_, err := collection.DeleteOne(*ctx, f)
	if err != nil {
		return err
	}

	return nil
}
