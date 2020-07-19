package app

import (
	"fmt"
	"net/http"
)

// StartApplication starts the application.
func StartApplication() {
	r := mapRoutes()

	// Starting the server.
	fmt.Println("The users server is listening!")
	http.ListenAndServe(":80", r)
}
