package main

import (
	"clean/internal/db/mongoDB"
	"clean/internal/service"
	httpTransport "clean/internal/transport/http"
	"fmt"
)

var forever chan bool

func Run() error {
	db := mongoDB.NewDatabase()
	s := service.NewService(db)
	h := httpTransport.NewHandler(s)

	err := h.Serve()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		return
	}

	<-forever
}
