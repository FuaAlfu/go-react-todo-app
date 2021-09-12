package main

import(
	"fmt"
	"log"
	"net/http"
	_"server/router"
)

func routerConfig(){
	r := router.Router()
	fmt.Println("servering on port 8081")

	log.Fatal(http.ListenAndServe(":8081",r))
}

func main(){
	routerConfig()
}