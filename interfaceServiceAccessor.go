package berlioz

/*
 * SERVICE ACCESSOR
 */
type ServiceAccessor struct {
	id    string
	peers PeerAccessor
}

// TBD
func (x ServiceAccessor) Endpoint(name string) ServiceEndpointAccessor {
	return ServiceEndpointAccessor{id: x.id, endpoint: name, peers: NewEndpointPeers(x.id, name)}
}

// TBD
func (x ServiceAccessor) All() map[string]interface{} {
	return x.peers.All()
}

// TBD
func (x ServiceAccessor) Get(identity string) interface{} {
	return x.peers.Get(identity)
}

// TBD
func (x ServiceAccessor) First() interface{} {
	return x.peers.First()
}

// TBD
func (x ServiceAccessor) Random() interface{} {
	return x.peers.Random()
}

// TBD
func (x ServiceAccessor) MonitorAll(callback func(map[string]interface{})) {
	x.peers.MonitorAll(callback)
}

// TBD
func (x ServiceAccessor) MonitorFirst(callback func(interface{})) {
	x.peers.MonitorFirst(callback)
}

// TBD
func (x ServiceAccessor) Request() PeerHttpRequester {
	return PeerHttpRequester{peers: x.peers}
}
