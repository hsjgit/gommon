package httputil

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func UploadFile(url, name, filename string, params url.Values, file io.Reader) ([]byte, error) {
	w := &bytes.Buffer{}
	writer := multipart.NewWriter(w)
	formFile, err := writer.CreateFormFile(name, filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	head := http.Header{}
	head.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	if params != nil {
		url = url + "?" + params.Encode()
	}
	req, err := http.NewRequest("POST", url, w)
	if err != nil {
		errors.Wrapf(err, "build req failed")
		return nil, err
	}
	req.Header = head
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "post failed")
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}

func UploadFiles(url string, file map[string][]io.Reader, params url.Values, field map[string][]string) ([]byte, error) {
	w := &bytes.Buffer{}
	writer := multipart.NewWriter(w)
	for s := range file {

		before, after, found := strings.Cut(s, "=")
		if !found {
			return nil, errors.New("file 参数格式错误，key必须为fieldname=filename")
		}

		for i := range file[s] {
			formFile, err := writer.CreateFormFile(before, after)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(formFile, file[s][i])
			if err != nil {
				return nil, err
			}
		}

	}
	for s := range field {
		for i := range field[s] {
			err := writer.WriteField(s, field[s][i])
			if err != nil {
				return nil, err
			}
		}

	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	head := http.Header{}
	head.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	if params != nil {
		url = url + "?" + params.Encode()
	}
	req, err := http.NewRequest("POST", url, w)
	if err != nil {
		errors.Wrapf(err, "build req failed")
		return nil, err
	}
	req.Header = head
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "post failed")
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}
