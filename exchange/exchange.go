package exchange

import (
	"github.com/pablonlr/arbibot/coin"
)

type Exchange interface {
	IsExchangeble(co coin.Coin) bool
	PriceUSD(co coin.Coin) (float64, error)
	GetExchangeAmount(amount float64, token1 coin.Coin, token2 coin.Coin) (float64, error)
	ID() string
}
