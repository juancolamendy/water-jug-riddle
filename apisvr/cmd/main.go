package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/juancolamendy/water-jug-riddle/apisvr/conf"
	"github.com/juancolamendy/water-jug-riddle/apisvr/route"
)

func main() {
	log.Println("Init apisvr")

    // Get application config
    config := conf.GetConfig()
    config.Dump()

    // Init routes
    rtr := route.GetRouter()

    // Create an HTTP server and start listening
    log.Printf("Start HTTP server. HTTP_PORT:[%s]\n", config.HTTP_PORT)
    http.ListenAndServe(fmt.Sprintf(":%s", config.HTTP_PORT),
        handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
        handlers.AllowedOrigins([]string{"*"}))(rtr))    
}