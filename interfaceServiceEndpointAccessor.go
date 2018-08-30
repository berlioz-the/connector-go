package berlioz

/*
 * SERVICE ENDPOINT ACCESSOR
 */
type ServiceEndpointAccessor struct {
	id       string
	endpoint string
	peers    PeerAccessor
}

// TBD
func (x ServiceEndpointAccessor) All() map[string]interface{} {
	return x.peers.All()
}

// TBD
func (x ServiceEndpointAccessor) Get(identity string) interface{} {
	return x.peers.Get(identity)
}

// TBD
func (x ServiceEndpointAccessor) First() interface{} {
	return x.peers.First()
}

// TBD
func (x ServiceEndpointAccessor) Random() interface{} {
	return x.peers.Random()
}

// TBD
func (x ServiceEndpointAccessor) MonitorAll(callback func(map[string]interface{})) SubscribeInfo {
	return x.peers.MonitorAll(callback)
}

// TBD
func (x ServiceEndpointAccessor) MonitorFirst(callback func(interface{})) SubscribeInfo {
	return x.peers.MonitorFirst(callback)
}

// TBD
func (x ServiceEndpointAccessor) Request() PeerHttpRequester {
	return PeerHttpRequester{peers: x.peers}
}
