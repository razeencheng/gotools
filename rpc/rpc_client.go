package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
}

var (
	DefaultClient = Client{&http.Client{Transport: http.DefaultTransport}}
	UserAgent     = "TrustAsia openApi golang client"
)

func newRequest(method, urlStr string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, urlStr, body)
	if err != nil {
		return
	}
	return
}

func (r Client) DoRequest(ctx context.Context, method, urlStr string) (resp *http.Response, err error) {
	req, err := newRequest(method, urlStr, nil)
	if err != nil {
		return
	}
	return r.Do(ctx, req)
}

func (r Client) DoRequestWith(ctx context.Context, method, urlStr, bodyType string, body io.Reader, bodyLength int) (resp *http.Response, err error) {
	req, err := newRequest(method, urlStr, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	req.ContentLength = int64(bodyLength)
	return r.Do(ctx, req)
}

func (r Client) DoRequestWith64(ctx context.Context, method, urlStr, bodyType string, body io.Reader, bodyLength int64) (resp *http.Response, err error) {
	req, err := newRequest(method, urlStr, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	req.ContentLength = bodyLength
	return r.Do(ctx, req)
}

func (r Client) DoRequestWithForm(ctx context.Context, method, urlStr string, data map[string][]string) (resp *http.Response, err error) {
	msg := url.Values(data).Encode()
	if method == "GET" || method == "HEAD" || method == "DELETE" {
		if strings.ContainsRune(urlStr, '?') {
			urlStr += "&"
		} else {
			urlStr += "?"
		}
		return r.DoRequest(ctx, method, urlStr+msg)
	}
	return r.DoRequestWith(ctx, method, urlStr, "application/x-www-form-urlencoded", strings.NewReader(msg), len(msg))
}

func (r Client) DoRequestWithJson(ctx context.Context, method, urlStr string, data interface{}) (resp *http.Response, err error) {
	msg, err := json.Marshal(data)
	if err != nil {
		return
	}
	return r.DoRequestWith(ctx, method, urlStr, "application/json", bytes.NewReader(msg), len(msg))
}

func (r Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", UserAgent)
	}

	transport := r.Transport // don't change r.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// avoid cancel() is called before Do(req), but isn't accurate
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	default:
	}

	if tr, ok := getRequestCanceler(transport); ok { // support CancelRequest
		reqC := make(chan bool, 1)
		go func() {
			resp, err = r.Client.Do(req)
			reqC <- true
		}()
		select {
		case <-reqC:
		case <-ctx.Done():
			tr.CancelRequest(req)
			<-reqC
			err = ctx.Err()
		}
	} else {
		resp, err = r.Client.Do(req)
	}
	return
}

type requestCanceler interface {
	CancelRequest(req *http.Request)
}

type nestedObjectGetter interface {
	NestedObject() interface{}
}

func getRequestCanceler(tp http.RoundTripper) (rc requestCanceler, ok bool) {

	if rc, ok = tp.(requestCanceler); ok {
		return
	}

	p := interface{}(tp)
	for {
		getter, ok1 := p.(nestedObjectGetter)
		if !ok1 {
			return
		}
		p = getter.NestedObject()
		if rc, ok = p.(requestCanceler); ok {
			return
		}
	}
}
