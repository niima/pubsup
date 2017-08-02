package models

//Subscriber is the service which is subscribed to
//a certain tag to be informed after each publish
type Subscriber struct {
	Tag    string
	Url    string
	Method string
}
