package main

import (
	"fmt"
	"net/http"

	"github.com/pablonlr/arbibot/exchange/cex/binance/pclient"
	/*
		"log"

		"github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/pablonlr/arbibot/coin"
		"github.com/pablonlr/arbibot/exchange/dex/bscdex"
	*/)

const address = "0x325E343f1dE602396E256B67eFd1F61C3A6B38Bd"

func main() {
	cl := pclient.Client{http.DefaultClient}
	resp, err := cl.Time()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	/*
		client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
		if err != nil {
			log.Fatal(err)
		}
		routeToUsd := []common.Address{
			common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
		}
		dex, err := bscdex.NewBSCDEX(address, client, routeToUsd)
		if err != nil {
			panic(err)
		}

		xlm := coin.NewBSCCoin("XLM", "Stellar", "0x43C934A845205F0b514417d757d7235B8f53f1B9")
		//usdt := coin.NewBSCCoin("USDT", "USDT", "0x55d398326f99059fF775485246999027B3197955")
		/*
			f, err := dex.GetExchangeAmount(1000, xlm, usdt)
			if err != nil {
				panic(err)
			}

		f := dex.IsExchangeble(xlm)

		println(f)
	*/

}
