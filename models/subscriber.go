package models

//Subscriber is the service which is subscribed to
//a certain tag to be informed after each publish
type Subscriber struct {
	Tag    string `json:"tag"`
	Chanel string `json:"chanel"`
	URL    string `json:"url"`
	Method string `json:"method"`
}

//SubscriberList is a list of subs
type SubscriberList struct {
	Collection []Subscriber
}
