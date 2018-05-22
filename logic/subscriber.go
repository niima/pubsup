package logic

import (
	"fmt"
	"strings"

	"git.raad.cloud/cloud/sync-service/models"
)

//ConsumerPool suppose to handle consumer processes and bind them to variuos http apis
func ConsumerPool() {
	subs := LoadConfig()
	for _, item := range subs {
		fmt.Println(item.Chanel)
		go PreCreateTopics(item)
		go NsqConsumer(item)

	}
}

//SubscriberHandler is the entrypoint for nsq, nsq need to call this method for
func SubscriberHandler(msg []byte, sub models.Subscriber) error {
	//entrypoint
	strMsg := string(msg)
	s := strings.Split(strMsg, "|$|^|")
	err := HTTPPost(sub.URL+s[0], sub.Method, []byte(s[1]))
	fmt.Println("posted to" + sub.URL)

	return err
}
