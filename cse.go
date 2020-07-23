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

// record holds the record data
type record struct {
	Title string
	Value float64
	Date  string
}

// market holds the market data in a specific date for historical market summary
type market struct {
	SL            int
	Date          string
	Trade         int64
	Volume        int64
	ValueInTK     float64
	MarketCapInMN float64
	CSE30         float64
	CSCX          float64
	CASPI         float64
	CSE50         float64
	CSI           float64
}

// Summary holds the historical market summaries array and the record trading or highest records data
type Summary struct {
	HighestRecords      []*record
	HistoricalSummaries []*market
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
	SL          int
	TradingCode string
	LTP         float64
	Open        float64
	High        float64
	Low         float64
	YCP         float64
	Trade       int64
	ValueInMN   float64
	Volume      int64
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
func (c *CSE) GetLatestPrices(by sortBy, order sortOrder) ([]*CSEShare, error) {
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

// GetMarketSummary returns the summary with highest records till now and the historical market summary data
func (c *CSE) GetMarketSummary() (*Summary, error) {
	summary := &Summary{
		HighestRecords:      nil,
		HistoricalSummaries: nil,
	}
	highestRecords := make([]*record, 0)
	historicalSummaries := make([]*market, 0)

	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/market/historical_market")
	if err != nil {
		return nil, errErrorFetchingUrl
	}

	// process to get the highest records data
	list := htmlquery.Find(doc, `//*[@id="wrapper"]/div/div/div[1]/div/div[3]/div[1]/div/div/div`)

	for _, v := range list {
		tabsCont := htmlquery.Find(v, "div")
		for i, v := range tabsCont {
			if i == 0 {
				continue
			}
			recordTitle := htmlquery.FindOne(v, `//*[@id="highscore_tab_1"]`)
			recordValue := htmlquery.FindOne(v, `//*[@id="highscore_tab_2"]`)
			recordDate := htmlquery.FindOne(v, `//*[@id="highscore_tab_3"]`)
			r := &record{
				Title: htmlquery.InnerText(recordTitle),
				Value: toFloat64(htmlquery.InnerText(recordValue)),
				Date:  htmlquery.InnerText(recordDate),
			}
			highestRecords = append(highestRecords, r)
		}
	}

	// process to get the historical market summary data
	list = htmlquery.Find(doc, `//*[@id="wrapper"]/div/div/div[1]/div/div[3]/div[2]/div/div/div`)

	for _, v := range list {
		tabsCont := htmlquery.Find(v, "div")
		for i, v := range tabsCont {
			if i == 0 {
				continue
			}

			historySL := htmlquery.FindOne(v, `//*[@id="market_tab_1"]`)
			historyDate := htmlquery.FindOne(v, `//*[@id="market_tab_2"]`)
			historyTrade := htmlquery.FindOne(v, `//*[@id="market_tab_3"]`)
			historyVolume := htmlquery.FindOne(v, `//*[@id="market_tab_4"]`)
			historyValueInTK := htmlquery.FindOne(v, `//*[@id="market_tab_5"]`)
			historyMarketCapInMN := htmlquery.FindOne(v, `//*[@id="market_tab_6"]`)
			historyCSE30 := htmlquery.FindOne(v, `//*[@id="market_tab_7"]`)
			historyCSCX := htmlquery.FindOne(v, `//*[@id="market_tab_8"]`)
			historyCASPI := htmlquery.FindOne(v, `//*[@id="market_tab_9"]`)
			historyCSE50 := htmlquery.FindOne(v, `//*[@id="market_tab_10"]`)
			historyCSI := htmlquery.FindOne(v, `//*[@id="market_tab_11"]`)

			m := &market{
				SL:            toInt(htmlquery.InnerText(historySL)),
				Date:          htmlquery.InnerText(historyDate),
				Trade:         toInt64(htmlquery.InnerText(historyTrade)),
				Volume:        toInt64(htmlquery.InnerText(historyVolume)),
				ValueInTK:     toFloat64(htmlquery.InnerText(historyValueInTK)),
				MarketCapInMN: toFloat64(htmlquery.InnerText(historyMarketCapInMN)),
				CSE30:         toFloat64(htmlquery.InnerText(historyCSE30)),
				CSCX:          toFloat64(htmlquery.InnerText(historyCSCX)),
				CASPI:         toFloat64(htmlquery.InnerText(historyCASPI)),
				CSE50:         toFloat64(htmlquery.InnerText(historyCSE50)),
				CSI:           toFloat64(htmlquery.InnerText(historyCSI)),
			}
			historicalSummaries = append(historicalSummaries, m)
		}
	}

	summary.HighestRecords = highestRecords
	summary.HistoricalSummaries = historicalSummaries
	return summary, nil
}
