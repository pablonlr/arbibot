package binance

import (
	"net/http"

	"github.com/pablonlr/arbibot/coin"
	"github.com/pablonlr/arbibot/exchange/cex/binance/pclient"
)

type Client struct {
	*pclient.Client
}

func NewBinanceCliente() *Client {
	return &Client{&pclient.Client{&http.Client{}}}
}

func (client *Client) IsExchangeble(co coin.Coin) bool {
	_, err := client.Client.AveragePrice(co.ID, "USDT")
	return err == nil
}
func (client *Client) PriceUSD(co coin.Coin) (float64, error) {
	bid, _, err := client.Client.UpdatedPrice(co.ID, "USDT")
	if err != nil {
		return 0, err
	}
	return bid, nil
}

func (client *Client) GetExchangeAmount(amount float64, token1 coin.Coin, token2 coin.Coin) (float64, error) {
	bid, _, err := client.Client.UpdatedPrice(token1.ID, token2.ID)
	if err == nil {
		return bid * amount, nil
	}
	_, ask, err := client.Client.UpdatedPrice(token2.ID, token1.ID)
	if err == nil {
		return amount / ask, nil
	}
	bid, _, err = client.Client.UpdatedPrice(token1.ID, "USDT")
	if err != nil {
		return 0, err
	}
	_, ask, err = client.Client.UpdatedPrice(token2.ID, "USDT")
	if err != nil {
		return 0, err
	}

	return 10 * bid / ask, nil
}

func (client *Client) ID() string {
	return "binance"
}
