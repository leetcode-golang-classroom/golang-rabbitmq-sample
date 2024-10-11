package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leetcode-golang-classroom/golang-rabbitmq-sample/internal"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to rabbitmq
	conn, err := amqp.Dial(internal.AppConfig.RABBITMQ_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// open a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	// subscribe the message to the queue
	messages, err := channel.Consume(internal.AppConfig.QUEUE_NAME, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case message := <-messages:
			log.Printf("Message: %s\n", message.Body)
		case <-signChan:
			log.Println("Interrupt detected")
			os.Exit(0)
		}
	}
}
