package berlioz

import ()

/*
 * MY ENDPOINT ACCESSOR
 */
type MyEndpointAccessor struct {
	name string
}

// TBD
func (x MyEndpointAccessor) Get() EndpointModel {
	path := make([]string, 1)
	// path[0] = x.name
	mapVal := registry.get("endpoints", path)
	if mapVal == nil {
		return EndpointModel{}
	}
	epMap := mapVal.(*map[string]EndpointModel)

	if ep, ok := (*epMap)[x.name]; ok {
		return ep
	}

	return EndpointModel{}
}

// TBD
func (x MyEndpointAccessor) Monitor(callback func(EndpointModel)) {
	currVal := x.Get()
	registry.subscribe("policies", nil, func(interface{}) {
		newVal := x.Get()
		if currVal != newVal {
			currVal = newVal
			callback(currVal)
		}
	})
	callback(currVal)
}
