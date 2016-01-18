package commands

import (
	"fmt"
	"net/http"
)

func FileServer(path string) {
	// deliver files from the directory /var/www
	//fileServer := http.FileServer(http.Dir("/var/www"))
	fileServer := http.FileServer(http.Dir(path))

	fmt.Println("Listening on localhost:8000")

	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)
	// That's it!
}
