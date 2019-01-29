package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

type MyHandler struct{}

func (m *MyHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("%s\n", string(msg.Body))
	return nil
}

func main() {
	nsqConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("mirza_topic", "x", nsqConfig)
	if err != nil {
		panic(err.Error())
	}

	consumer.AddHandler(&MyHandler{})

	if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		panic(err.Error())
	}

	<-consumer.StopChan

}
