package main

import (
	"net/http"
	"os"
)

func main() {

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)

	http.HandleFunc("/cgi-bin/printenv", printEnv)

	http.HandleFunc("/hello", hello)
	
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

}

func hello(writer http.ResponseWriter, req *http.Request) {

writer.Write([]byte("<h1>Hello</h1>"))

}

func printEnv(writer http.ResponseWriter, req *http.Request) {

	env := os.Environ()

	writer.Write([]byte("<h1>Enviroment</h1>\n<pre>"))

	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}

	writer.Write([]byte("</pre>"))

}
