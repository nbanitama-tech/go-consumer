package main

import (
	nsq "github.com/nsqio/go-nsq"
	"encoding/json"
	"log"
)

type message1Handler struct{}
type message1 struct {
 	Name   		string
	Address 	string
}

const 
(
	Consumer1 = "consumer-1"
)

func initConsumer1()ConsumerClient{
	return ConsumerClient{
		Name: Consumer1,
		Topic: "topic",
		Channel: "channel",
		Handler: &message1Handler{},
	}
}

func (h *message1Handler) HandleMessage(m *nsq.Message) error {
	var request message1
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Println(err)
	  	return err
	}

	go request.printMessage()

	// set as finish
	return nil
}

func (m *message1) printMessage(){
	log.Printf("Incoming message (%+v)\n", m)
}