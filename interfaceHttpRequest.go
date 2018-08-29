package berlioz

import (
	"context"
	"encoding/json"
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
type PeerHttpRequester struct {
	peers NewPeerAccessor
}

// TBD
func (x PeerHttpRequester) Get(ctx context.Context, url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(ctx, req)
}

// TBD
func (x PeerHttpRequester) Post(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(ctx, req)
}

// TBD
func (x PeerHttpRequester) PostForm(ctx context.Context, url string, data url.Values) (*http.Response, []byte, error) {
	return x.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

// TBD
func (x PeerHttpRequester) Delete(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return x.Do(ctx, req)
}

// TBD
func (x PeerHttpRequester) Head(ctx context.Context, url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return x.Do(ctx, req)
}

// TBD
func (x PeerHttpRequester) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	f := func(rawPeer interface{}, span *TracingSpan) ([]interface{}, error) {

		byteData, _ := json.Marshal(rawPeer)
		peer := EndpointModel{}
		err := json.Unmarshal(byteData, &peer)
		if err != nil {
			log.Printf("[TEST] Could not convert peer.")
			return nil, fmt.Errorf("Invalid peer provided.")
		}

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

	res, err := execute(ctx, x.peers, req.Method, f)
	if err != nil {
		return nil, nil, err
	}
	return res[0].(*http.Response), res[1].([]byte), nil
}
