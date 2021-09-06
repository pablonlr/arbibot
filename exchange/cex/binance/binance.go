package binance

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func Test() {
	a := binance.NewClient("a", "b")
	fmt.Println(a)
}
