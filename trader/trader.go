package trader

import (
	"github.com/pablonlr/arbibot/coin"
	"github.com/pablonlr/arbibot/exchange"
)

var (
	suscribExchanges []exchange.Exchange

	//Coin ID is assosiated with a list of exchanges for arbitrage
	exToTrade map[string][]exchange.Exchange
)

func init() {
	suscribExchanges = []exchange.Exchange{}
	exToTrade = make(map[string][]exchange.Exchange)
}

func SuscribeCoin(co coin.Coin, tradeableIn ...exchange.Exchange) {
	tradeableIn = deleteRepeatedExAndSuscribe(tradeableIn)
	exToTrade[co.ID] = append(exToTrade[co.ID], tradeableIn...)
}

func deleteRepeatedExAndSuscribe(exs []exchange.Exchange) []exchange.Exchange {
	var brk bool
	for i, x := range exs {
		for _, y := range suscribExchanges {
			if x.ID() == y.ID() {
				exs = append(exs[:i], exs[i+1:]...)
				brk = true
				break
			}
		}
		if !brk {
			suscribExchanges = append(suscribExchanges, x)
		}
		brk = false
	}
	return exs
}

func GetExchangeAmountsOfCoin(amount float64, coin, referenceCoin coin.Coin) (map[string]float64, error) {
	prices := make(map[string]float64)
	exs := exToTrade[coin.ID]
	for _, x := range exs {
		p, err := x.GetExchangeAmount(amount, coin, referenceCoin)
		if err != nil {
			return nil, err
		}
		prices[x.ID()] = p
	}
	return prices, nil
}
