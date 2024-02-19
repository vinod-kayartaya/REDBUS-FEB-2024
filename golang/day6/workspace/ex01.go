package main

import (
	"fmt"
	"net/http"
)

func home(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, `<html><head><title>Go Server
	</title></head><body bgcolor="lightblue"><h1>Hello, world!</h1>
	<hr />
	<p>This is a simple output from a go server</p>
	</body></html>`)
}

func main() {
	fmt.Println("starting the http server at port 7788")
	http.HandleFunc("/", home)
	http.ListenAndServe("0.0.0.0:7788", nil)
}
