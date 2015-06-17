package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 80, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)
	
	log.Printf("Running on port %d\n", *port)
	
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}