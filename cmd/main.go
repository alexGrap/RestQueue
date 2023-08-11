package main

import (
	"context"
	"inter/internal/handlers"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	countOfRoutine, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal()
	}
	ctx, ctxCancel := context.WithCancel(context.Background())
	closedChanel := make(chan bool, 1)
	server := handlers.Fabric(countOfRoutine, ctx, closedChanel)
	go func() {
		err := server.Hearing()
		if err != nil {

		}
	}()
	shotDown := make(chan os.Signal, 1)
	signal.Notify(shotDown, syscall.SIGINT, syscall.SIGTERM)
	<-shotDown
	ctxCancel()
	countOfClosed := 0

	//creating the "while" loop to collect messages about finishing all goRoutines
	for {
		value := <-closedChanel

		if value {
			countOfClosed += 1
		}
		if countOfClosed == 1 {
			break
		}
	}
	close(closedChanel)

	log.Println("Program end")

}
