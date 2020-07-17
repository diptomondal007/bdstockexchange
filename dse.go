package bd_sotck_exchange

import (
	"github.com/antchfx/htmlquery"
	"log"
)

type dse struct {
}

func NewDSE() *dse {
	return new(dse)
}

type latest struct {
	ID          string `json:"id"`
	TradingCode string `json:"trading_code"`
	LTP         string `json:"ltp"`
	High        string `json:"high"`
	Low         string `json:"low"`
	CloseP      string `json:"close_p"`
	YCP         string `json:"ycp"`
	Change      string `json:"change"`
	Trade       string `json:"trade"`
	Value       string `json:"value"`
	Volume      string `json:"volume"`
}

func (d *dse) GetLatestUpdate() []*latest {
	latestShares := make([]*latest, 0)
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/latest_share_price_all.php")
	if err != nil {
		log.Println(err)
	}
	list, err := htmlquery.QueryAll(doc, "//table //tr")
	if err != nil {
		panic(err)
	}
	//log.Println(len(list))
	for i, n := range list {
		//log.Println(a)
		if i == 0 {
			continue
		}
		a := htmlquery.Find(n, "//td")
		s := &latest{
			ID:          htmlquery.InnerText(a[0]),
			TradingCode: htmlquery.InnerText(a[1]),
			LTP:         htmlquery.InnerText(a[2]),
			High:        htmlquery.InnerText(a[3]),
			Low:         htmlquery.InnerText(a[4]),
			CloseP:      htmlquery.InnerText(a[5]),
			YCP:         htmlquery.InnerText(a[6]),
			Change:      htmlquery.InnerText(a[7]),
			Trade:       htmlquery.InnerText(a[8]),
			Value:       htmlquery.InnerText(a[9]),
			Volume:      htmlquery.InnerText(a[10]),
		}
		latestShares = append(latestShares, s)
	}
	return latestShares
}
