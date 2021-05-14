package main

import (
	nsq "github.com/nsqio/go-nsq" 
	"log"
)

var config *nsq.Config
	
type MessageHandler interface {
	HandleMessage(m *nsq.Message) error
}

type ConsumerClient struct{
	Name string
	Topic string
	Channel string
	Handler MessageHandler
}

type ConsumerHub struct{
	consumers map[string]*nsq.Consumer
	Register chan ConsumerClient
	Stop chan string
}

func InitHub(c *nsq.Config)ConsumerHub{
	config = c
	return ConsumerHub{
		consumers: make(map[string]*nsq.Consumer),
		Register: make(chan ConsumerClient),
		Stop: make(chan string),
	}
}

func (c *ConsumerHub) Run(){
	for {
		select {
		case consumer, ok := <-c.Register:
			if !ok {
				log.Println("Unable to register!!!!")
			} else {
				err := c.register(consumer)
				if err != nil {
					log.Println("Unable to create consumer!!!!")
				}
			}
		case name, ok := <- c.Stop:
			if !ok {
				log.Println("Unable to stop!!!!")
			} else {
				c.stop(name)
			}
		}
	}
}

func (c *ConsumerHub) register(data ConsumerClient) error {
	con, err := nsq.NewConsumer(data.Topic, data.Channel, config)
	if err != nil {
		return err
	} 
	con.AddHandler(data.Handler)

	//nsqlookupd is the domain
	con.ConnectToNSQLookupd("nsqlookupd:4161")	
	c.consumers[data.Name] = con
	log.Printf("Consumer %s registered!\n", data.Name)

	return nil
}

func (c *ConsumerHub) stop(name string){
	con := c.consumers[name]
	con.Stop()
	delete(c.consumers, name)
	log.Printf("Consumer %s stopped!!\n", name)
}