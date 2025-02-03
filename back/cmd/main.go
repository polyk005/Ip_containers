package main

import (
	"fmt"
	"log"

	"github.com/polyk005/Ip_containers/back"
)

func main() {
	fmt.Print("SERVER Run")
	srv := new(back.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
