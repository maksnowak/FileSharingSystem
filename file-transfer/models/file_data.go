package models

type FileData struct {
	Path string `json:"path"`
	Data []byte `json:"data"`
}

type FileDataResponse struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

type FileDataRequest struct {
	Path string `json:"path"`
}
