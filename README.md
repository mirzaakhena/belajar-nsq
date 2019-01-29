# belajar-nsq

This program require nsqd to run

please refer to https://nsq.io/overview/quick_start.html

There are 3 different program

## consumer
This only run the consumer

## producer
This only run the producer

## restapi
This run both consumer and producer. Producer send the value via rest api

```GET 127.0.0.1:8080/publish?message=Hello```

Then consumer can see receive a message via console
