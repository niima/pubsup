package logic

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.pec.ir/cloud/sync-service/models"

	"github.com/bitly/go-nsq"
)

//NsqProducer produce message for nsq
func NsqProducer(topic string, body []byte) error {
	nsqdIP := "127.0.0.1"
	if os.Getenv("NSQD_IP") != "" {
		nsqdIP = os.Getenv("NSQD_IP")
	}
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqdIP+":4150", config)

	err := w.Publish(topic, body)
	if err != nil {
		log.Println("nima: could not connect" + nsqdIP)
		log.Println("nima: " + err.Error())

	}

	w.Stop()
	return err
}

//NsqConsumer is a method for consume messages from nsq
func NsqConsumer(sub models.Subscriber) {

	lookupIP := "127.0.0.1"
	if os.Getenv("LOOKUP_IP") != "" {
		lookupIP = os.Getenv("LOOKUP_IP")
	}

	var reader *nsq.Consumer
	var err error
	lookup := lookupIP + ":4161"
	conf := nsq.NewConfig()
	conf.Set("maxInFlight", 1000)
	reader, err = nsq.NewConsumer(sub.Tag, sub.Chanel, conf)
	fmt.Println("new consumer " + sub.Chanel + "/" + sub.Tag)
	if err != nil {
		fmt.Println(err)
	}

	reader.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		//output of this handler
		fmt.Println("adding nsq handler")
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
