package models

// ONLY FOR SWAGGER DOCS

type HTTP200 struct {
	Message string `json:"message" example:"Operation successful"`
}

type HTTP400 struct {
	Message string `json:"message" example:"Invalid request body"`
}

type HTTP404 struct {
	Message string `json:"message" example:"Could not find requested data"`
}

type HTTP500 struct {
	Message string `json:"message" example:"Error while processing request"`
}

type Register struct {
	Username     string `json:"username" example:"Karol_Wojtyla"`
	Email        string `json:"email" example:"huan.pablo.dos@vatican.city"`
	PasswordHash string `json:"passwordHash" example:"Kremowki"`
	PasswordSalt string `json:"passwordSalt" example:"Slony_Karmel"`
	Role         string `json:"role" example:"admin"`
}

type Update struct {
	Email        string   `json:"email" example:"huan.pablo.tres@vatican.city"`
	PasswordHash string   `json:"passwordHash" example:"Papiezowki"`
	PasswordSalt string   `json:"passwordSalt" example:"Pozdrawiam_Polakow"`
	OwnedFiles   []string `json:"ownedFiles" example:"rower,pies,zachrystia"`
	SharedFiles  []string `json:"sharedFiles" example:"zaba,cialo_chrystusa"`
}

type Salt struct {
	Username     string `json:"username" example:"Karol_Wojtyla"`
	PasswordSalt string `json:"passwordSalt" example:"Slony_Karmel"`
}
