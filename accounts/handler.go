package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var counter int = 0

func addOne(x *int) int {
	*x += 1
	return *x
}

// Hello godoc
//
//	@Description	Increment the counter and greet the user
//	@Produce		json
//	@Success		200	{string}	string "Example: Hello from 'accounts' times 3!"
//	@Router			/hello [get]
func Hello(w http.ResponseWriter, _ *http.Request) {
	msg := "Hello from 'accounts' times " + strconv.Itoa(addOne(&counter)) + "!"
	jsonStr, _ := json.Marshal(msg)
	_, _ = fmt.Fprintf(w, "%s\n", jsonStr)
}
