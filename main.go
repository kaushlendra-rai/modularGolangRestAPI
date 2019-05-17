package main

import (
	"log"
	"net/http"

	"com.kaushal/rai/handler"
	"com.kaushal/rai/util"
)

func main() {
	router := handler.Initialize()
	name := util.GetProperty("application.name")
	log.Println("application.name: ", name)
	log.Fatal(http.ListenAndServe(":80", router))
}
