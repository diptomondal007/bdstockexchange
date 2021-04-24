package main

import (
	"log"

	"github.com/diptomondal007/bdstockexchange"
)

func main() {
	dse := bdstockexchange.NewDSE()
	ms, err := dse.GetMarketSummary()
	if err != nil {
		log.Println(err)
	}

	log.Println(ms.DseX.DSEXIndex)
}
