package http

import (
	"errors"
	_ "fmt"
	"goworker/log"
	"io/ioutil"
	_http "net/http"
	_url "net/url"
	"strings"
)

var (
	Verbose bool

	ErrInvalidUrl = errors.New("Invalid URL")
)

func init() {
	Verbose = true
}

type HttpMethod string

const (
	Post = HttpMethod("POST")
	Get  = HttpMethod("GET")
	Put  = HttpMethod("Put")
)

func HttpGet(url string) (string, error) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return string(""), ErrInvalidUrl
	}

	resp, err := _http.Get(url)
	if err != nil {
		// TODO handle error
		log.Error(err)
		return string(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO handle error
		log.Error(err)
		return string(""), err
	}

	return string(body), nil
}

func HttpPost(url string, data string) (string, error) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return string(""), ErrInvalidUrl
	}
	resp, err := _http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		// TODO handler error
		log.Error(err)
		return string(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO handle error
		return string(""), err
	}

	return string(body), nil
}

func HttpPostForm(url string, data *map[string]string) (string, error) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return string(""), ErrInvalidUrl
	}
	var values _url.Values
	values = _url.Values{}
	if data != nil {
		if len(*data) > 0 {
			for k, v := range *data {
				values.Add(k, v)
			}
		}
	}

	if values == nil {
		return string(""), ErrInvalidUrl
	}

	resp, err := _http.PostForm(url, values)

	if err != nil {
		// TODO handle error
		return string(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO handle error
		return string(""), err
	}

	return string(body), nil
}

func HttpDo(url string, method HttpMethod, data string, header *map[string]string) (string, error) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		return string(""), ErrInvalidUrl
	}
	client := &_http.Client{}

	req, err := _http.NewRequest(string(method), url, strings.NewReader(data))
	if err != nil {
		// TODO handle error
		return string(""), err
	}

	if header != nil {
		if len(*header) > 0 {
			for k, v := range *header {
				req.Header.Set(k, v)
			}
		}
	}

	if method == Post {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	// TODO cookie
	//req.Header.Set("Cookie", "yyy=xxx;zzz=sss")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO handle error
		return string(""), err
	}

	return string(body), nil
}
