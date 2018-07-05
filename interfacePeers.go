package berlioz

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
)

// TBD
type PeerAccessor struct {
	kind string
	path []string
}

// TBD
func Peers(kind string, name string, endpoint string) PeerAccessor {
	path := make([]string, 2)
	path[0] = name
	path[1] = endpoint
	return PeerAccessor{kind: kind, path: path}
}

func (x PeerAccessor) getMap() indexedMap {
	return registry.getAsIndexedMap(x.kind, x.path)
}

// TBD
func (x PeerAccessor) Monitor(callback func(PeerAccessor)) {
	registry.subscribe(x.kind, x.path, func(interface{}) {
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
	f := func(rawPeer interface{}, span *TracingSpan) ([]interface{}, error) {
		peer := rawPeer.(EndpointModel)

		req.URL.Scheme = peer.Protocol
		req.URL.Host = peer.Address + ":" + strconv.Itoa(int(peer.Port))

		if span != nil {
			tracer := (*span).tracer
			innerSpan := (*span).span
			if tracer != nil && innerSpan != nil {
				if err := (*tracer).Inject(
					(*innerSpan).Context(),
					opentracing.TextMap,
					opentracing.HTTPHeadersCarrier(req.Header),
				); err != nil {
					fmt.Printf("error encountered while trying to inject span: %+v\n", err)
				}
			}
		}

		log.Printf("Request: %s", req.URL.String())
		resp, err := new(http.Client).Do(req)
		if err != nil {
			// fmt.Printf("Response: %s, error: %s\n", resp, err)
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// fmt.Printf("Response: %s, error: %s\n", resp, err)
			return nil, err
		}

		return []interface{}{resp, body}, nil
	}

	res, err := execute(ctx, x.kind, x.path, f)
	if err != nil {
		return nil, nil, err
	}
	return res[0].(*http.Response), res[1].([]byte), nil
}
