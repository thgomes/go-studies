package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

var i int64
var c1 = make(chan Message)
var c2 = make(chan Message)

func rcvMsgFromRabbitMQ() {
	for {
		time.Sleep(time.Second)
		atomic.AddInt64(&i, 1)
		msg := Message{i, "Hello from RabbitMQ"}
		c1 <- msg
	}
}

func rcvMsgFromKafka() {
	for {
		time.Sleep(time.Second)
		atomic.AddInt64(&i, 1)
		msg := Message{i, "Hello from RabbitMQ"}
		c2 <- msg
	}
}

func main() {
	go rcvMsgFromRabbitMQ()
	go rcvMsgFromKafka()
	for {
		select {
		case msg1 := <-c1: // rabbitmq
			fmt.Printf("recived ID: %d - Msg: %s\n", msg1.Id, msg1.Msg)
		case msg2 := <-c2: // kafka
			fmt.Printf("recived ID: %d - Msg: %s\n", msg2.Id, msg2.Msg)
		}
	}
}
