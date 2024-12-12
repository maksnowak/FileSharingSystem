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
