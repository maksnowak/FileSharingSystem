package db

import (
	"context"

	"file-transfer/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateFile(ctx *context.Context, collection *mongo.Collection, f models.File) (models.File, error) {
	res, err := collection.InsertOne(*ctx, f)
	if err != nil {
		return f, err
	}

	f.FileID = res.InsertedID.(bson.ObjectID)
	return f, nil
}

func GetAllFiles(ctx *context.Context, collection *mongo.Collection) ([]models.File, error) {
	var files []models.File
	cursor, err := collection.Find(*ctx, bson.D{})
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
	filter := bson.M{"_id": f.FileID}

	res := collection.FindOne(*ctx, filter)
	err := res.Decode(&f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func GetFilesByUserID(ctx *context.Context, collection *mongo.Collection, userID string) ([]models.File, error) {
	var files []models.File
	filter := bson.M{"userID": userID}

	cursor, err := collection.Find(*ctx, filter)
	if err != nil {
		return files, err
	}

	err = cursor.All(*ctx, &files)
	if err != nil {
		return files, err
	}

	return files, nil
}

func UpdateFile(ctx *context.Context, collection *mongo.Collection, f models.File) (models.File, error) {
	filter := bson.M{"_id": f.FileID}

	update := bson.M{
		"$set": bson.M{
			"fileName":  f.FileName,
			"userID":    f.UserID,
			"tags":      f.Tags,
			"path":      f.Path,
			"blobURL":   f.BlobURL,
			"hasAccess": f.HasAccess,
		},
	}

	res := collection.FindOneAndUpdate(*ctx, filter, update)
	err := res.Decode(&f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func DeleteFile(ctx *context.Context, collection *mongo.Collection, f models.File) error {
	_, err := collection.DeleteOne(*ctx, f)
	if err != nil {
		return err
	}

	return nil
}
