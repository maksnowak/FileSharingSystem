package app

import (
	"net/http"
)

func (a *App) initRoutes() {
	a.Router.HandleFunc("/health", a.healthCheck).Methods(http.MethodGet)

	a.Router.HandleFunc("/file", a.createFile).Methods(http.MethodPost)
	a.Router.HandleFunc("/files", a.getAllFiles).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.getFile).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.updateFile).Methods(http.MethodPut)
	a.Router.HandleFunc("/file/{file_id}", a.deleteFile).Methods(http.MethodDelete)
}
