package main

import (
	"file-transfer/app"
)

//	@title			File transfer API
//	@version		0.2
//	@description	Webserver providing saving and retrieval of files from MongoDB

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath	/
func main() {
	a := App{}
	a.Initialize()

	a.Run(":8080")
}
