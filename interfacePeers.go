package berlioz

/*
 * GLOBAL ACCESSORS
 */

/*
 * NEW Peer Accessor
 */
type PeerAccessor struct {
	serviceID string
	path      []string
}

// TBD
func NewEndpointPeers(id string, endpoint string) PeerAccessor {
	path := make([]string, 2)
	path[0] = id
	path[1] = endpoint
	return PeerAccessor{serviceID: id, path: path}
}

// TBD
func NewResourcePeers(name string) PeerAccessor {
	path := make([]string, 1)
	path[0] = name
	return PeerAccessor{path: path}
}

func (x PeerAccessor) getMap() IndexedMap {
	return registry.getAsIndexedMap("peer", x.path)
}

// TBD
func (x PeerAccessor) All() map[string]interface{} {
	y := x.getMap()
	return y.all()
}

// TBD
func (x PeerAccessor) Get(identity string) interface{} {
	y := x.getMap()
	return y.get(identity)
}

// TBD
func (x PeerAccessor) First() interface{} {
	y := x.getMap()
	return y.first()
}

// TBD
func (x PeerAccessor) Random() interface{} {
	y := x.getMap()
	return y.random()
}

// TBD
func (x PeerAccessor) MonitorAll(callback func(map[string]interface{})) SubscribeInfo {
	return registry.subscribe("peer", x.path, func(value interface{}) {
		callback(value.(IndexedMap).all())
	})
}

// TBD
func (x PeerAccessor) MonitorFirst(callback func(interface{})) SubscribeInfo {
	return x.monitorPeer(firstKeySelector, callback)
}

type peerSelectorT func(IndexedMap) interface{}

// TBD
func (x PeerAccessor) monitorPeer(selector peerSelectorT, callback func(interface{})) SubscribeInfo {
	// oldValue := interface{}(nil)
	return registry.subscribe("peer", x.path, func(peers interface{}) {
		value := selector(peers.(IndexedMap))
		callback(value)
	})
}
