package berlioz

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TBD
type PeerAccessor struct {
	kind string
	path []string
}

// TBD
func Peers(kind string, name string, endpoint string) PeerAccessor {
	var path []string
	path = make([]string, 2)
	path[0] = name
	path[1] = endpoint
	return PeerAccessor{kind: kind, path: path}
}

func (x PeerAccessor) getMap() indexedMap {
	return registry.getAsIndexedMap(x.kind, x.path)
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
func (x PeerAccessor) Get(identity string) EndpointModel {
	y := x.getMap()
	return y.get(identity).(EndpointModel)
}

// TBD
func (x PeerAccessor) Random() EndpointModel {
	y := x.getMap()
	return y.random().(EndpointModel)
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
func (x PeerRequester) Get(url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(req)
}

// TBD
func (x PeerRequester) Post(url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(req)
}

// TBD
func (x PeerRequester) PostForm(url string, data url.Values) (*http.Response, []byte, error) {
	return x.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

// TBD
func (x PeerRequester) Delete(url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(req)
}

// TBD
func (x PeerRequester) Head(url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(req)
}

// TBD
func (x PeerRequester) Do(req *http.Request) (*http.Response, []byte, error) {
	f := func(peer interface{}) ([]interface{}, error) {
		ep := peer.(EndpointModel)

		req.URL.Scheme = ep.Protocol
		req.URL.Host = ep.Address + ":" + strconv.Itoa(int(ep.Port))

		// TODO: Manipulate headers.

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

	res, err := execute(x.kind, x.path, f)
	if err != nil {
		return nil, nil, err
	}
	return res[0].(*http.Response), res[1].([]byte), nil
}
