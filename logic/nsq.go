package logic

import (
	"errors"
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

//Producer produce message for nsq
func Producer(topic string, body []byte) {

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish(topic, body)
	if err != nil {
		log.Panic("could not connect")

	}

	w.Stop()

}

//Consumer is a method for consume messages from nsq
func Consumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("event", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		if len(message.Body) == 0 {
			// returning an error results in the message being re-enqueued
			// a REQ is sent to nsqd
			return errors.New("body is blank re-enqueue message")
		}

		// Let's log our message!
		//log.Print(message.Body)

		log.Printf("got a message: %v", string(message.Body))
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("could not connect")
	}
	wg.Wait()
}

func start() {
	//go
}
