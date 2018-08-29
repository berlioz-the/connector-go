package berlioz

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
 * GLOBAL ACCESSORS
 */

/*
 * Interface Peer Accessor
 */
type IPeerAccessor interface {
	All() PeersModel
}

/*
 * NEW Peer Accessor
 */
type NewPeerAccessor struct {
	serviceID string
	path      []string
}

// TBD
func NewEndpointPeers(id string, endpoint string) NewPeerAccessor {
	path := make([]string, 2)
	path[0] = id
	path[1] = endpoint
	return NewPeerAccessor{serviceID: id, path: path}
}

// TBD
func NewResourcePeers(name string) NewPeerAccessor {
	path := make([]string, 1)
	path[0] = name
	return NewPeerAccessor{path: path}
}

func (x NewPeerAccessor) getMap() indexedMap {
	return registry.getAsIndexedMap("peer", x.path)
}

// TBD
func (x NewPeerAccessor) All() map[string]interface{} {
	y := x.getMap()
	return y.all()
}

// TBD
func (x NewPeerAccessor) Get(identity string) interface{} {
	y := x.getMap()
	return y.get(identity)
}

// TBD
func (x NewPeerAccessor) First() interface{} {
	y := x.getMap()
	return y.first()
}

// TBD
func (x NewPeerAccessor) Random() interface{} {
	y := x.getMap()
	return y.random()
}

// TBD
func (x NewPeerAccessor) MonitorAll(callback func(indexedMap)) {
	registry.subscribe("peer", x.path, func(value interface{}) {
		callback(value.(indexedMap))
	})
}

// TBD
func (x NewPeerAccessor) MonitorFirst(callback func(interface{})) {
	x.monitorPeer(firstKeySelector, callback)
}

type peerSelectorT func(indexedMap) interface{}

// TBD
func (x NewPeerAccessor) monitorPeer(selector peerSelectorT, callback func(interface{})) {
	// oldValue := interface{}(nil)
	registry.subscribe("peer", x.path, func(peers interface{}) {
		value := selector(peers.(indexedMap))
		callback(value)
	})
}

/************************* LEGACY ********************/
// TBD
type PeerAccessor struct {
	path []string
}

// TBD
func Peers(kind string, name string, endpoint string) PeerAccessor {
	path := make([]string, 2)
	path[0] = name
	path[1] = endpoint
	return PeerAccessor{path: path}
}

func (x PeerAccessor) getMap() indexedMap {
	return registry.getAsIndexedMap("peer", x.path)
}

// TBD
func (x PeerAccessor) Monitor(callback func(PeerAccessor)) {
	registry.subscribe("peer", x.path, func(interface{}) {
		callback(x)
	})
}

// TBD
func (x PeerAccessor) All() PeersModel {
	y := x.getMap()
	result := PeersModel{}
	for k, v := range y.all() {
		result[k] = v.(EndpointModel)
	}
	return result
}

// TBD
func (x PeerAccessor) Get(identity string) (EndpointModel, bool) {
	y := x.getMap()
	val := y.get(identity)
	if val == nil {
		return EndpointModel{}, false
	}
	return val.(EndpointModel), true
}

// TBD
func (x PeerAccessor) Random() (EndpointModel, bool) {
	y := x.getMap()
	val := y.random()
	if val == nil {
		return EndpointModel{}, false
	}
	return val.(EndpointModel), true
}

// TBD
type PeerRequester struct {
	kind string
	path []string
}

func (x PeerRequester) getMap() indexedMap {
	return registry.getAsIndexedMap(x.kind, x.path)
}

// TBD
func Request(kind string, name string, endpoint string) PeerRequester {
	var path []string
	path = make([]string, 2)
	path[0] = name
	path[1] = endpoint
	return PeerRequester{kind: kind, path: path}
}

// TBD
func (x PeerRequester) Get(ctx context.Context, url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(ctx, req)
}

// TBD
func (x PeerRequester) Post(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(ctx, req)
}

// TBD
func (x PeerRequester) PostForm(ctx context.Context, url string, data url.Values) (*http.Response, []byte, error) {
	return x.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

// TBD
func (x PeerRequester) Delete(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(ctx, req)
}

// TBD
func (x PeerRequester) Head(ctx context.Context, url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(ctx, req)
}

// TBD
func (x PeerRequester) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	return nil, nil, nil
}
