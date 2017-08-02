package logic

//SubscriberHandler is the entrypoint for nsq, nsq need to call this method for
func SubscriberHandler(msg []byte, chanel string) {
	//entrypoint for
	var url string
	switch chanel {
	case "social-service":
		url = "http://192.168.95.84:5005"
	case "messaging-service":
		url = "http://192.168.95.84:5006"
	case "search-service":
		url = "http://192.168.95.84:5007"
	}

	HTTPPost(url, msg)
}
