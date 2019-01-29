package main

import nsq "github.com/nsqio/go-nsq"

func main() {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		panic(err.Error())
	}
	producer.Publish("mirza_topic", []byte("Helloo"))

}
