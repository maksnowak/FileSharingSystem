type File struct {
	FileID    primitive.ObjectID `json:"id" "bson:"_id"`
	FileName  string             `json:"file_name" "bson:"fileName"`
	UserID    string             `json:"user_id" bson:"userID"`
	Tags      []string           `json:"tags" bson:"tags"`
	Data      []byte             `json:"data" bson:"data"`
	HasAccess []string           `json:"has_access" "bson:"hasAccess"` // List of user IDs
}
