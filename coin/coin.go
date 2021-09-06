package coin

import (
	"github.com/ethereum/go-ethereum/common"
)

type Coin struct {
	ID                string
	Name              string
	ContractAddresses map[string]common.Address
}

func NewBSCCoin(id, name, contractAddr string) Coin {
	return Coin{
		ID:                id,
		Name:              name,
		ContractAddresses: map[string]common.Address{"bsc": common.HexToAddress(contractAddr)},
	}
}
