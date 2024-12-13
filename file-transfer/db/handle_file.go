package db

import (
	"context"

	"file-transfer/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateFile(ctx *context.Context, collection *mongo.Collection, f models.File) error {
	_, err := collection.InsertOne(*ctx, f)
	if err != nil {
		return err
	}

	return nil
}

func GetAllFiles(ctx *context.Context, collection *mongo.Collection) ([]models.File, error) {
	var files []models.File
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

func GetFile(ctx *context.Context, collection *mongo.Collection, f models.File) (models.File, error) {
	err := collection.FindOne(*ctx, f).Decode(f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func UpdateFile(ctx *context.Context, collection *mongo.Collection, f models.File) error {
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

func DeleteFile(ctx *context.Context, collection *mongo.Collection, f models.File) error {
	_, err := collection.DeleteOne(*ctx, f)
	if err != nil {
		return err
	}

	return nil
}
