package pclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
}

const baseURL = "https://api3.binance.com"

//GetReq Make a GET request following the endopoints received
func (client *Client) GetReq(endpoints ...string) ([]byte, error) {
	endpoints = append([]string{"/api/v3"}, endpoints...)
	url := baseURL + strings.Join(endpoints, "/")
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return client.doReq(req)
}
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
	if resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}
	return body, nil
}
