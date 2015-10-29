package main

import (
	"flag"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	filename := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()
	fmt.Println(*filename)
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
