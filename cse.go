package bdstockexchange

import (
	"errors"
	"log"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
)

// CSE is a struct to access cse related methods
type CSE struct {
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

//NewCSE returns new CSE object
func NewCSE() *CSE {
	return new(CSE)
}

// CSEShare is a model for a single company's latest price data provided by the cse website
type CSEShare struct {
	SL          int     `json:"id"`
	TradingCode string  `json:"trading_code"`
	LTP         float64 `json:"ltp"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	YCP         float64 `json:"ycp"`
	Trade       int64   `json:"trade"`
	ValueInMN   float64 `json:"value"`
	Volume      int64   `json:"volume"`
}

func getCSELatestPrices() ([]*CSEShare, error) {
	shares := make([]*CSEShare, 0)

	doc, err := htmlquery.LoadURL("https://www.CSE.com.bd/market/current_price")
	if err != nil {
		log.Fatal(err)
	}

	list := htmlquery.Find(doc, `//*[@id="dataTable"]/tbody/tr`)

	for _, v := range list {
		td := htmlquery.Find(v, "//td")
		s := &CSEShare{}
		for i, t := range td {
			switch i {
			case slCSE:
				s.SL = toInt(htmlquery.InnerText(t))
				break
			case stockCodeCSE:
				s.TradingCode = strings.TrimSpace(htmlquery.InnerText(t))
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
	return shares, nil
}

// GetLatestPrices returns the array of latest share prices or error in case of any error
// It takes by which field the array should be sorted ex: SortByTradingCode and sort order ex: ASC
// It will return an error for if user tries to sort with a non existing file in the CSEShare model or invalid category name or invalid sort order
func (c *CSEShare) GetLatestPrices(by sortBy, order sortOrder) ([]*CSEShare, error) {
	arr, err := getCSELatestPrices()
	if err != nil {
		return nil, err
	}
	return arr, err
}

func sortCse(arr []*CSEShare, by sortBy, order sortOrder) ([]*CSEShare, error) {
	if order != ASC && order != DESC {
		return nil, errors.New("order param is not valid. put a ASC or DESC as order param")
	}
	switch by {
	case SortByTradingCode:
		if order == ASC {
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].TradingCode > arr[j].TradingCode
		})
		return arr, nil
	case SortByHighPrice:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].High < arr[j].High
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].High > arr[j].High
		})
		return arr, nil
	case SortByLTP:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].LTP < arr[j].LTP
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].LTP > arr[j].LTP
		})
		return arr, nil
	case SortByOpeningPrice:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].Open < arr[j].Open
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Open > arr[j].Open
		})
		return arr, nil
	case SortByLowPrice:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].Low < arr[j].Low
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Low > arr[j].Low
		})
		return arr, nil
	case SortByNumberOfTrades:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].Trade < arr[j].Trade
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Trade > arr[j].Trade
		})
		return arr, nil
	case SortByValue:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].ValueInMN < arr[j].ValueInMN
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].ValueInMN > arr[j].ValueInMN
		})
		return arr, nil
	case SortByVolumeOfShare:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].Volume < arr[j].Volume
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Volume > arr[j].Volume
		})
		return arr, nil
	case SortByYCP:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].YCP < arr[j].YCP
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].YCP > arr[j].YCP
		})
		return arr, nil
	default:
		return nil, errors.New("sorting with the given sort by param is not possible. try another one")
	}
}
