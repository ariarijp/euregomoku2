package commands

import (
	"net/http"
	"os"
)

func PrintEnv(path string) {
	// file handler for most files
	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", fileServer)

	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printEnv)

	// deliver requests to the handlers
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
	// That's it!
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}
