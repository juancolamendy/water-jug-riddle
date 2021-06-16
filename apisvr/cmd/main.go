package main

import (
	"log"

	"github.com/juancolamendy/water-jug-riddle/apisvr/conf"
)

func main() {
	log.Println("Init apisvr")

    // Get application config
    config := conf.GetConfig()
    config.Dump()	
}