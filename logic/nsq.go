package logic

import (
	"fmt"
	"log"

	"github.com/bitly/go-nsq"
)

//NsqProducer produce message for nsq
func NsqProducer(topic string, body []byte) {

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish(topic, body)
	if err != nil {
		log.Panic("could not connect")

	}

	w.Stop()

}

func loop(inChan chan *nsq.Message, chanel string) {
	for msg := range inChan {
		SubscriberHandler(msg.Body, chanel)
		fmt.Println(string(msg.Body) + "\n FROM CHANEL:" + chanel)
		msg.Finish()
	}
}

//NsqConsumer is a method for consume messages from nsq
func NsqConsumer(topic string, chanel string) {
	var reader *nsq.Consumer
	var err error
	inChan := make(chan *nsq.Message)
	lookup := "127.0.0.1:4161"
	conf := nsq.NewConfig()
	conf.Set("maxInFlight", 1000)
	reader, err = nsq.NewConsumer(topic, chanel, conf)
	if err != nil {
		fmt.Println(err)
	}
	reader.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		inChan <- m
		return nil
	}))
	err = reader.ConnectToNSQLookupd(lookup)
	if err != nil {
		fmt.Println(err)
	}
	go loop(inChan, chanel)
	<-reader.StopChan
}

func ConsumerPool() {
	go NsqConsumer("user", "social-service")
	go NsqConsumer("user", "messaging-service")
	go NsqConsumer("product", "search-service")
	go NsqConsumer("product", "social-service")
}
