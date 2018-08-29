package berlioz

// TBD
type NativeResourceAccessor struct {
	id    string
	peers PeerAccessor
}

// TBD
func (x NativeResourceAccessor) All() map[string]interface{} {
	return x.peers.All()
}

// TBD
func (x NativeResourceAccessor) Get(identity string) interface{} {
	return x.peers.Get(identity)
}

// TBD
func (x NativeResourceAccessor) First() interface{} {
	return x.peers.First()
}

// TBD
func (x NativeResourceAccessor) Random() interface{} {
	return x.peers.Random()
}

// TBD
func (x NativeResourceAccessor) MonitorAll(callback func(map[string]interface{})) {
	x.peers.MonitorAll(callback)
}

// TBD
func (x NativeResourceAccessor) MonitorFirst(callback func(interface{})) {
	x.peers.MonitorFirst(callback)
}
