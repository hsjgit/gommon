package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func JsonPost(addr string, req interface{}, resp interface{}) (err error) {
	var headers = http.Header{}
	headers.Add(contentTypeKey, ContentTypeJson)
	headers.Add("Accept", "application/json")

	var reqBody = &bytes.Buffer{}
	enc := json.NewEncoder(reqBody)
	if err = enc.Encode(req); err != nil {
		return errors.Wrap(err, "JsonPost encode failed:")
	}

	var respData []byte
	respData, err = HttpPost(addr, reqBody, nil, headers)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respData, resp)
	return
}
