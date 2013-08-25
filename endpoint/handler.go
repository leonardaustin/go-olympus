package endpoint

import ()

func List() map[int]string {

	endpointtest := map[int]string{
		123: "endpoint1",
		234: "endpoint2",
		345: "endpoint3",
		456: "endpoint4",
	}

	return endpointtest
}

func Register(ep *Endpoint) {
	// Add to the list of active endpoints
}
