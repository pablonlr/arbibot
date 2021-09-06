package pclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
}

const baseURL = "https://api3.binance.com"

func (client *Client) doReq(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//GetReq Make a GET request following the endopoints received and the req parameters in body
func (client *Client) GetReq(params map[string]string, endpoints ...string) ([]byte, error) {
	endpoints = append([]string{"/api/v3"}, endpoints...)
	ur := baseURL + strings.Join(endpoints, "/")
	req, err := http.NewRequest("GET", ur, nil)
	if err != nil {
		return nil, err
	}
	queries := url.Values{}
	for x := range params {
		queries.Add(x, params[x])
	}
	req.URL.RawQuery = queries.Encode()
	fmt.Println((ur))
	return client.doReq(req)
}
