package endpoint

import (
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	"log"
	"strconv"
)

var (
	endpoints map[string]*Endpoint
)

func init() {
	// Need to run make in method
	endpoints = make(map[string]*Endpoint)
}

func List() map[string]*Endpoint {

	return endpoints
}

func Register(ep *Endpoint) {
	endpoints[ep.Name+ep.Hostname+strconv.Itoa(ep.ProcessId)] = ep
	log.Println("New Endpoint Register : ", ep.Name+ep.Hostname+strconv.Itoa(ep.ProcessId))
}

func TestRegister() {

	testendpoint := Endpoint{
		Name:      "testName",
		ProcessId: 29383,
		Hostname:  "hostname",
		AuthLevel: 5.8,
		Handler:   Testfunc,
	}

	fmt.Println(testendpoint)
	Register(&testendpoint)
}

func Testfunc(req *Request) (proto.Message, error) {
	var m proto.Message
	err := errors.New("test error")
	return m, err

}
