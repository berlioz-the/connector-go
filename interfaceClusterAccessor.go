package berlioz

/*
 * SERVICE ACCESSOR
 */
type ClusterAccessor struct {
	id    string
	peers PeerAccessor
}

// TBD
func (x ClusterAccessor) Endpoint(name string) ServiceEndpointAccessor {
	return ServiceEndpointAccessor{id: x.id, endpoint: name, peers: NewEndpointPeers(x.id, name)}
}

// TBD
func (x ClusterAccessor) All() map[string]interface{} {
	return x.peers.All()
}

// TBD
func (x ClusterAccessor) Get(identity string) interface{} {
	return x.peers.Get(identity)
}

// TBD
func (x ClusterAccessor) First() interface{} {
	return x.peers.First()
}

// TBD
func (x ClusterAccessor) Random() interface{} {
	return x.peers.Random()
}

// TBD
func (x ClusterAccessor) MonitorAll(callback func(map[string]interface{})) SubscribeInfo {
	return x.peers.MonitorAll(callback)
}

// TBD
func (x ClusterAccessor) MonitorFirst(callback func(interface{})) SubscribeInfo {
	return x.peers.MonitorFirst(callback)
}

// TBD
func (x ClusterAccessor) Request() PeerHttpRequester {
	return PeerHttpRequester{peers: x.peers}
}
