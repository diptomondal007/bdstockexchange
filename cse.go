package bdstockexchange

import (
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
)

type cse struct {
}

const (
	slCSE = iota
	stockCodeCSE
	ltpCSE
	openCSE
	highCSE
	lowCSE
	ycpCSE
	tradeCSE
	valueCSE
	volumeCSE
)

func NewCSE() *cse {
	return new(cse)
}

type latestShares []*cseShare

type cseShare struct {
	SL        int     `json:"id"`
	STOCKCode string  `json:"trading_code"`
	LTP       float64 `json:"ltp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	YCP       float64 `json:"ycp"`
	Trade     int64   `json:"trade"`
	ValueInMN float64 `json:"value"`
	Volume    int64   `json:"volume"`
}


func getCSELatestPrices() ([]*cseShare, error) {
	shares := make([]*cseShare, 0)

	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/market/current_price")
	if err != nil {
		log.Fatal(err)
	}

	list := htmlquery.Find(doc, `//*[@id="dataTable"]/tbody/tr`)

	for _, v := range list {
		td := htmlquery.Find(v, "//td")
		s := &cseShare{}
		for i, t := range td {
			switch i {
			case slCSE:
				s.SL = toInt(htmlquery.InnerText(t))
				break
			case stockCodeCSE:
				s.STOCKCode = strings.TrimSpace(htmlquery.InnerText(t))
				break
			case ltpCSE:
				s.LTP = toFloat64(htmlquery.InnerText(t))
				break
			case openCSE:
				s.LTP = toFloat64(htmlquery.InnerText(t))
				break
			case highCSE:
				s.High = toFloat64(htmlquery.InnerText(t))
				break
			case lowCSE:
				s.Low = toFloat64(htmlquery.InnerText(t))
				break
			case ycpCSE:
				s.YCP = toFloat64(htmlquery.InnerText(t))
				break
			case tradeCSE:
				s.Trade = toInt64(htmlquery.InnerText(t))
				break
			case valueCSE:
				s.ValueInMN = toFloat64(htmlquery.InnerText(t))
				break
			case volumeCSE:
				s.Volume = toInt64(htmlquery.InnerText(t))
				break
			}
		}
		shares = append(shares, s)
	}

	for _, v := range shares {
		log.Println(v.STOCKCode)
	}
	return shares, nil
}
