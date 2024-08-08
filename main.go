package main

import (
	"fmt"

	besttimetobuyandsellstock "github.com/banggibima/go-sea-lb-prepare/best-time-to-buy-and-sell-stock"
	jumpgame "github.com/banggibima/go-sea-lb-prepare/jump-game"
	missingnumber "github.com/banggibima/go-sea-lb-prepare/missing-number"
)

func main() {
	fmt.Println(besttimetobuyandsellstock.MaxProfit([]int{8, 4, 3, 2, 1}))

	fmt.Println(jumpgame.CanReach([]int{2, 3, 1, 1, 4}))

	fmt.Println(missingnumber.FindMissingNumber([]int{3, 0, 1}))
}
