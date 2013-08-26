package endpoint

import (
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	"log"
	"strconv"
)

var (
	activeEndpoints map[string][]*Endpoint
)

func init() {
	// Need to run make in method
	activeEndpoints = make(map[string][]*Endpoint)
}

func List() map[string][]*Endpoint {

	return activeEndpoints
}

func Register(ep *Endpoint) {

	currentEndpoints := activeEndpoints[ep.Name]
	newEndpoint := []*Endpoint{ep}

	// Need ... as newEndpoint is not a slice
	activeEndpoints[ep.Name] = append(currentEndpoints, newEndpoint...)
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

	fmt.Println("Create New Endpoint", testendpoint)
	Register(&testendpoint)

	testendpoint2 := Endpoint{
		Name:      "testName",
		ProcessId: 9999,
		Hostname:  "hostname2",
		AuthLevel: 2.8,
		Handler:   Testfunc,
	}

	fmt.Println("Create New Endpoint", testendpoint2)
	Register(&testendpoint2)

	testendpoint3 := Endpoint{
		Name:      "testName2",
		ProcessId: 9999,
		Hostname:  "hostname",
		AuthLevel: 0.8,
		Handler:   Testfunc,
	}

	fmt.Println("Create New Endpoint", testendpoint3)
	Register(&testendpoint3)
}

func Testfunc(req *Request) (proto.Message, error) {
	var m proto.Message
	err := errors.New("test error")
	return m, err

}
