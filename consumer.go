package main

import (
	"log"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
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
