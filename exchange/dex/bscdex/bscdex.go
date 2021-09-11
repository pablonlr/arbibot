package bscdex

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pablonlr/arbibot/coin"
	"github.com/pablonlr/arbibot/exchange/dex/bscdex/router"
)

const (
	RPCserver           string = "https://bsc-dataseed.binance.org/"
	networkID           string = "bsc"
	BUSDContractAddress string = "0xe9e7cea3dedca5984780bafc599bd69add087d56"
	precision           int    = 100000
)

type BSCDEX struct {
	RouterContract string
	Client         *ethclient.Client
	routerInstance *router.Router
	pathToUSDPrice []common.Address
	id             string
}

func NewBSCDEX(id, routerAddress string, client *ethclient.Client, pathToUSD []common.Address) (*BSCDEX, error) {
	addr := common.HexToAddress(routerAddress)
	instance, err := router.NewRouter(addr, client)
	if err != nil {
		return nil, err
	}
	return &BSCDEX{
		RouterContract: routerAddress,
		Client:         client,
		routerInstance: instance,
		pathToUSDPrice: pathToUSD,
		id:             id,
	}, nil

}

func (dex *BSCDEX) IsExchangeble(co coin.Coin) bool {
	res, err := dex.PriceUSD(co)
	if err != nil {
		return false
	}
	return res > 0

}
func (dex *BSCDEX) ID() string {
	return dex.id
}
func (dex *BSCDEX) PriceUSD(co coin.Coin) (float64, error) {
	path := append([]common.Address{co.ContractAddresses[networkID]}, dex.pathToUSDPrice...)
	result, err := dex.getExchangeAmount(precision, path)
	if err != nil {
		return 0, err
	}
	return float64(result) / float64(precision), nil

}

func (dex *BSCDEX) GetExchangeAmount(amount float64, token1 coin.Coin, token2 coin.Coin) (float64, error) {
	route := []common.Address{
		token1.ContractAddresses[networkID],
		token2.ContractAddresses[networkID],
	}
	return dex.getExchangeAmount(int(amount), route)

}

func (dex *BSCDEX) getExchangeAmount(amount int, path []common.Address) (float64, error) {
	bg := big.NewInt(int64(amount))
	result, err := dex.routerInstance.GetAmountsOut(&bind.CallOpts{}, bg, path)
	if err != nil {
		return 0, err
	}
	return float64(result[len(result)-1].Int64()), nil

}
