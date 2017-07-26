package logic

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

func Producer(topic string, body []byte) {

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish(topic, body)
	if err != nil {
		log.Panic("could not connect")

	}

	w.Stop()

}

func Consumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("event", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("got a message: %v", message)
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
