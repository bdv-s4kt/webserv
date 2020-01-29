package handler

import (
	"fmt"
	"net/http"
)

// SayHello is a simple handler which print message on server side and return simple HTML as response
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is server side message")
	fmt.Fprintln(w, `<b>Response</b> from <i>server!</i>`)
}
