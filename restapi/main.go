package main

import (
	"fmt"
	"net/http"

	"github.com/nsqio/go-nsq"

	"github.com/gin-gonic/gin"
)

type MyStruct struct {
}

func (m *MyStruct) HandleMessage(msg *nsq.Message) error {
	fmt.Printf(">>> %s\n", string(msg.Body))
	return nil
}

func main() {

	topic := "mirza_topic"
	addr := "127.0.0.1:4150"
	nsqConfig := nsq.NewConfig()

	// prepare producer
	producer, err := nsq.NewProducer(addr, nsqConfig)
	if err != nil {
		panic(err.Error())
	}

	// prepare consumer (in go routine)
	go func() {
		consumer, err := nsq.NewConsumer(topic, "ch", nsqConfig)
		if err != nil {
			panic(err.Error())
		}
		consumer.AddHandler(&MyStruct{})

		// start to listen
		if err := consumer.ConnectToNSQD(addr); err != nil {
			panic(err.Error())
		}
		<-consumer.StopChan
	}()

	// prepare rest service api
	router := gin.Default()
	router.GET("/publish", func(c *gin.Context) {
		message := c.DefaultQuery("message", "")
		if message == "" {
			return
		}
		producer.Publish(topic, []byte(message))
		c.JSON(http.StatusOK, map[string]string{"message": message})
	})

	// run it
	router.Run()
}
