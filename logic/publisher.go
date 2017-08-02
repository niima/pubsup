//Package logic :
// - add publisher
// - remove publisher
// - publish
package logic

//Publish action logic for publish controller
// - this logic gets data from controller and pass it to nsq
func Publish(tag string, data string) {
	//log it
	//produce nsq message
	NsqProducer(tag, []byte(data))

}
