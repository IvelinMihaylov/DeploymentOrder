package main

import (
	"fmt"
	"log"
	"net/http"
	"HandlingDeploymentOrder/src/web/RestControllers"
)

func main() {
	http.HandleFunc("/", RestControllers.OrderedComps)
	fmt.Println("Server working...\nServer listening on port \"8081\".")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
