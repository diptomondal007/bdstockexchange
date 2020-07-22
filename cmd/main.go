package main

import (
	"log"

	"github.com/diptomondal007/bdstockexchange"
)

func main() {
	dse := bdstockexchange.NewDSE()
	arr, err := dse.GetLatestPrices(bdstockexchange.SortByVolumeOfShare, bdstockexchange.DESC)
	if err != nil {
		log.Println(err)
	}
	for _, v := range arr {
		log.Println(v.TradingCode, v.Volume)
	}
	//c := bdStockExchange.NewCSE()
}
