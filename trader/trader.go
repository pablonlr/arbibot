package trader

import (
	"errors"

	"github.com/pablonlr/arbibot/coin"
	"github.com/pablonlr/arbibot/exchange"
	"github.com/pablonlr/arbibot/log"
)

type Ticker interface {
	Tick(trad *Trader)
}

type Arbitrader struct {
	Trader
	Tic Ticker
}

type Trader struct {
	logger           log.Logger
	suscribExchanges []exchange.Exchange
	//Coin ID is assosiated with a list of exchanges for arbitrage
	exToTrade map[*coin.Coin][]exchange.Exchange
}

func NewTrader(log log.Logger) *Trader {
	return &Trader{
		logger:           log,
		suscribExchanges: []exchange.Exchange{},
		exToTrade:        make(map[*coin.Coin][]exchange.Exchange),
	}
}

func (trad *Trader) GetLogger() *log.Logger {
	return &trad.logger
}

func (trad *Trader) SuscribeCoin(co coin.Coin, tradeableIn ...exchange.Exchange) {
	tradeableIn = trad.deleteRepeatedExAndSuscribe(tradeableIn)
	trad.exToTrade[&co] = append(trad.exToTrade[&co], tradeableIn...)
}

func (trad *Trader) deleteRepeatedExAndSuscribe(exs []exchange.Exchange) []exchange.Exchange {
	var brk bool
	for i, x := range exs {
		for _, y := range trad.suscribExchanges {
			if x.ID() == y.ID() {
				exs = append(exs[:i], exs[i+1:]...)
				brk = true
				break
			}
		}
		if !brk {
			trad.suscribExchanges = append(trad.suscribExchanges, x)
		}
		brk = false
	}
	return exs
}

func (trad *Trader) GetExchangeAmountsOfCoin(amount float64, coin, referenceCoin coin.Coin) (map[string]float64, error) {
	prices := make(map[string]float64)
	exs, ok := trad.exToTrade[&coin]
	if !ok {
		return nil, errors.New("coin is not registered")
	}
	for _, x := range exs {
		p, err := x.GetExchangeAmount(amount, coin, referenceCoin)
		if err != nil {
			return nil, err
		}
		prices[x.ID()] = p
	}
	return prices, nil
}

func (trad *Trader) GetPricesInUSD(coin *coin.Coin) (map[string]float64, error) {
	prices := make(map[string]float64)
	exs, ok := trad.exToTrade[coin]
	if !ok {
		return nil, errors.New("coin is not registered")
	}
	for _, x := range exs {
		p, err := x.PriceUSD(*coin)
		if err != nil {
			return nil, err
		}
		prices[x.ID()] = p
	}
	return prices, nil
}
