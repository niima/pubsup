package logic

import "gitlab.pec.ir/cloud/sync-service/models"

//ConsumerPool suppose to handle consumer processes and bind them to variuos http apis
func ConsumerPool() {

	consumers := []models.Subscriber{
		{
			Tag:    "user",
			Chanel: "search-service",
			URL:    "http://192.168.95.84:5007",
			Method: "",
		},
		{
			Tag:    "user",
			Chanel: "social-service",
			URL:    "http://192.168.95.84:5005",
			Method: "",
		},
		{
			Tag:    "product",
			Chanel: "social-service",
			URL:    "http://192.168.95.84:5005",
			Method: "",
		},
		{
			Tag:    "product",
			Chanel: "messaging-service",
			URL:    "http://192.168.95.84:5006",
			Method: "",
		},
	}

	for _, item := range consumers {
		go NsqConsumer(&item)

	}
}

//SubscriberHandler is the entrypoint for nsq, nsq need to call this method for
func SubscriberHandler(msg []byte, sub *models.Subscriber) error {
	//entrypoint
	err := HTTPPost(sub.URL, msg)
	return err
}
