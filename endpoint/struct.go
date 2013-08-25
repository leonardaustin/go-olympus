package endpoint

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/bitly/nsq/nsq"
)

// Request for all handlers
type Request struct {
	Message nsq.Message
}

// All end points require a name, authlevel (set in config) and handler method
type Endpoint struct {
	Name      string
	AuthLevel float64
	Handler   Handler
}

type Handler func(req *Request) (proto.Message, error)
