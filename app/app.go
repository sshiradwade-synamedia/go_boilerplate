package app

import (
	"fmt"
	"go_boilerplate/app/router"
	"log"
	"net/http"
)

func Run() {
	r := router.NewRouter()
	const PORT = "8000"
	fmt.Println("Server Listening on port: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
