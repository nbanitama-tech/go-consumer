package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var hub ConsumerHub

func main() {
	log.Println("Starting the service")
	
	c := nsq.NewConfig()
	c.MaxAttempts = 1
	c.MaxInFlight = 10
	c.MaxRequeueDelay = time.Second * 60
	c.DefaultRequeueDelay = time.Second * 0

	hub = InitHub(c)
	go hub.Run()
	log.Println("Hub started!!!")

	newConsumer := initConsumer1()
	hub.Register <- newConsumer

	// hub.Stop <- newConsumer.Name

   	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("The service terminated!!")
}