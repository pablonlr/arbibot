package trader

import (
	"fmt"
	"math"

	"github.com/pablonlr/arbibot/log"
)

type PricesInUsdTicker struct {
	priceDiverPerCent float64
}

func NewPricesInUsdTicker(priceDivergentPerCent float64) PricesInUsdTicker {
	return PricesInUsdTicker{
		priceDiverPerCent: priceDivergentPerCent,
	}
}

func (ti PricesInUsdTicker) Tick(trad *Trader) {
	for co := range trad.exToTrade {
		mp, err := trad.GetPricesInUSD(co)
		if err != nil {
			trad.logger.Log(fmt.Sprintf("Error getting price of %s \nError description: %s", co.Name, err.Error()))
		}
		comparePricewithDiference(mp, co.Name, ti.priceDiverPerCent, trad.logger)

	}
}

func comparePricewithDiference(prices map[string]float64, coinName string, priceDiverPerCent float64, inform log.Logger) {
	max := 0.
	var exMax, exMin string
	min := math.MaxFloat64
	for ex, pr := range prices {
		inform.Log(fmt.Sprintf("%s price in %s is %f", coinName, ex, pr))
		if pr > max {
			max = pr
			exMax = ex
		} else if pr < min {
			min = pr
			exMin = ex
		}
	}
	if max != 0 && min != math.MaxFloat64 && (max*(1-(priceDiverPerCent/100)) > min) {
		inform.Notify(fmt.Sprintf("Arbitrage oportunity for %s: \nPrice in %s is %f\nPrice in %s is %f\nDifference: %f%  ", coinName, exMin, min, exMax, max, 1-(max/min)))
	}

}
