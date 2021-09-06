package pclient

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (c *Client) Time() (string, error) {
	resp, err := c.GetReq(nil, "time")
	if err != nil {
		return "", err
	}
	return string(resp), nil

}

func (c *Client) AveragePrice(token1, token2 string) (string, error) {
	market := strings.ToUpper(token1 + token2)
	params := make(map[string]string)
	params["symbol"] = market
	resp, err := c.GetReq(params, "avgPrice")
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

type OrderBook struct {
	LastUpdateID int
	Bids         [][]string
	Asks         [][]string
}

func (c *Client) OrderBook(token1, token2 string, limit int) (*OrderBook, error) {
	market := strings.ToUpper(token1 + token2)
	params := make(map[string]string)
	params["symbol"] = market
	params["limit"] = strconv.Itoa(limit)
	resp, err := c.GetReq(params, "depth")
	if err != nil {
		return nil, err
	}
	book := &OrderBook{}
	err = json.Unmarshal(resp, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (c *Client) UpdatedPrice(token1, token2 string) (bid float64, ask float64, err error) {
	book, err := c.OrderBook(token1, token2, 5)
	if err != nil {
		return 0, 0, err
	}
	bid, err = strconv.ParseFloat((*book).Bids[0][0], 64)
	if err != nil {
		return 0, 0, err
	}
	ask, err = strconv.ParseFloat((*book).Asks[0][0], 64)
	if err != nil {
		return 0, 0, err
	}
	return
}
