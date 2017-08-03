package logic

import (
	"fmt"
	"log"
	"time"

	"gitlab.pec.ir/cloud/sync-service/models"

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

//NsqConsumer is a method for consume messages from nsq
func NsqConsumer(sub *models.Subscriber) {
	var reader *nsq.Consumer
	var err error
	lookup := "127.0.0.1:4161"
	conf := nsq.NewConfig()
	conf.Set("maxInFlight", 1000)
	reader, err = nsq.NewConsumer(sub.Tag, sub.Chanel, conf)
	if err != nil {
		fmt.Println(err)
	}

	reader.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		//output of this handler
		err := SubscriberHandler(msg.Body, sub)
		if err != nil {
			msg.Requeue(10 * time.Second)
			fmt.Println("MSG REQued")
		} else {
			fmt.Println(string(msg.Body) + "\n")
			msg.Finish()
		}
		return err
	}))
	err = reader.ConnectToNSQLookupd(lookup)
	if err != nil {
		fmt.Println(err)
	}
	<-reader.StopChan
}
