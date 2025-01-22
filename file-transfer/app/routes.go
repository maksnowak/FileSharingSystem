package app

import (
	_ "file-transfer/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func (a *App) initRoutes() {
	a.Router.HandleFunc("/health", a.healthCheck).Methods(http.MethodGet)

	a.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	log.Println("Swagger available at: http://localhost:8080/swagger/index.html")

	a.Router.HandleFunc("/file", a.createFile).Methods(http.MethodPost)
	a.Router.HandleFunc("/files", a.getAllFiles).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.getFile).Methods(http.MethodGet)
	a.Router.HandleFunc("/file/{file_id}", a.updateFile).Methods(http.MethodPut)
	a.Router.HandleFunc("/file/{file_id}", a.deleteFile).Methods(http.MethodDelete)
	a.Router.HandleFunc("/files/user/{user_id}", a.getFilesByUser).Methods(http.MethodGet)

	a.Router.HandleFunc("file/upload", a.uploadFile).Methods(http.MethodPost)
	a.Router.HandleFunc("file/download", a.downloadFile).Methods(http.MethodGet)
}
