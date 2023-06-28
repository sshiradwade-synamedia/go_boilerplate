package controller

import (
	"fmt"
	service "go_boilerplate/app/service/hello"
	"net/http"
)

type HelloHandler struct {
	CustomeMessage string
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg, err := service.Hello()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, msg)

}
