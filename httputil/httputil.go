package httputil

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type Options struct {
	Params url.Values
	Header http.Header
	Body   io.Reader
	Method string
	URL    string
}

func HttpGet(url string, params url.Values, headers *http.Header) ([]byte, error) {
	var empty = []byte{}
	if params != nil && len(params) != 0 {
		url = url + "?" + params.Encode()
	}
	req, _ := http.NewRequest("GET", url, nil)

	if headers != nil && len(*headers) != 0 {
		req.Header = *headers
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		return empty, err
	}
	for k := range resp.Header {
		headers.Set(k, resp.Header.Get(k))
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	return data, nil
}

func HttpPost(dest string, body io.Reader, params url.Values, headers *http.Header) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second,
	}

	if params != nil {
		dest = dest + "?" + params.Encode()
	}

	req, err := http.NewRequest("POST", dest, body)
	if err != nil {
		errors.Wrapf(err, "build req failed")
		return nil, err
	}

	if headers != nil {
		req.Header = *headers
	} else {
		req.Header.Add("Content-Type", "text/plain")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "post failed")
	}
	for k := range resp.Header {
		headers.Set(k, resp.Header.Get(k))
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	return data, nil
}

func HttpRequest(opt Options) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	dest := opt.URL
	if opt.Params != nil {
		dest = dest + "?" + opt.Params.Encode()
	}
	req, err := http.NewRequest(opt.Method, dest, opt.Body)
	if err != nil {
		errors.Wrapf(err, "build req failed")
		return nil, err
	}

	if opt.Header != nil {
		req.Header = opt.Header
	} else {
		req.Header.Add("Content-Type", "application/octet-stream")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "request failed")
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	return data, nil
}
