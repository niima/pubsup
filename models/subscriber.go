package models

//Subscriber is the service who is subscribed to
//special tag to be informed after each publish
type Subscriber struct {
	Tag    string
	URL    string
	Method string
}
