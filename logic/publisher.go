//Package logic :
// - add publisher
// - remove publisher
// - publish
package logic

//Publish action logic for publish controller
// - this logic gets data from controller and pass it to nsq
func Publish(tag string, urlParam string, data string) error {
	//log it
	//produce nsq message
	return NsqProducer(tag, []byte(urlParam+"|$|^|"+data))
}
