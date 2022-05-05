package main

import (
	"clean/internal/db/mongoDB"
	Logger "clean/internal/log"
	"clean/internal/service"
	httpTransport "clean/internal/transport/http"
	"fmt"
)

var forever chan bool

func Run() error {
	log := Logger.NewLogger()
	db := mongoDB.NewDatabase(log)
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
