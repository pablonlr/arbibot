package trader

import (
	"fmt"
	"time"
)

type TickerController struct {
	Tic                Ticker
	tickTimeMiliSecond int
	trd                *Trader
}

func (ctr *TickerController) StartTicker() {
	go func() {
		for now := range time.Tick(time.Millisecond * time.Duration(ctr.tickTimeMiliSecond)) {
			ctr.Tic.Tick(ctr.trd)
			ctr.trd.logger.Log(fmt.Sprintln("Checked at:", now))

		}
	}()
}

func NewTickerController(tic Ticker, tickMiliSeconds int, trad *Trader) *TickerController {
	return &TickerController{
		Tic:                tic,
		tickTimeMiliSecond: tickMiliSeconds,
		trd:                trad,
	}
}
