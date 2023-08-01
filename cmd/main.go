package main

import (
	"inter/internal/handlers"
	"inter/internal/usecase"
	"log"
	"os"
	"strconv"
)

func main() {

	usecase.CountOfGoingRoutine, _ = strconv.Atoi(os.Args[1])
	server := handlers.Fabric()
	err := server.Hearing()
	if err != nil {
		log.Fatal(err)
	}

}
