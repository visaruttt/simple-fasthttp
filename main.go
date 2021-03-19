package main

import (
	"flag"
	"github.com/valyala/fasthttp"
	"github.com/visaruttt/simple-fasthttp/router"
	"log"
)

func main() {

	port := "8000"
	addr := flag.String("addr", ":"+port, "Running on port")

	r := router.New()
	router.Mount(r)

	log.Printf("API Service is running on port %v\n", *addr)
	if err := fasthttp.ListenAndServe(*addr, r.Handler); err != nil {
		log.Fatal(err)
	}

}
