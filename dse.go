package bdstockexchange

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

// DSE is a struct to access dse related methods
type DSE struct {
}

const (
	idDSE = iota
	tradingCodeDSE
	ltpDSE
	highDSE
	lowDSE
	closePriceDSE
	ycpDSE
	changeDSE
	tradeDSE
	valueDSE
	volumeDSE
)

// NewDSE returns new DSE object
func NewDSE() *DSE {
	return new(DSE)
}

// DSEShare is a model for a single company's latest price data provided by the dse website
type DSEShare struct {
	ID          int     `json:"id"`
	TradingCode string  `json:"trading_code"`
	LTP         float64 `json:"ltp"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	CloseP      float64 `json:"close_p"`
	YCP         float64 `json:"ycp"`
	Change      float64 `json:"change"`
	Trade       int64   `json:"trade"`
	ValueInMN   float64 `json:"value"`
	Volume      int64   `json:"volume"`
}

func getDSELatestPrices(url string) ([]*DSEShare, error) {
	latestShares := make([]*DSEShare, 0)
	// Request the HTML page.
	if url == "" {
		url = "https://www.dsebd.org/latest_share_price_scroll_l.php"
	}
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/latest_share_price_scroll_l.php")
	if err != nil {
		return nil, err
	}
	tab, _ := htmlquery.QueryAll(doc, `/html/body/div[2]/section/div/div[3]/div[1]/div[2]/div[1]`)
	for _, v := range tab {
		tbody, err := htmlquery.QueryAll(v, "//tbody")
		if err != nil {
			return nil, err
		}
		for _, v := range tbody {
			td, err := htmlquery.QueryAll(v, "//tr //td")
			if err != nil {
				return nil, err
			}
			s := &DSEShare{}
			for index, v := range td {
				switch index {
				case idDSE:
					s.ID = toInt(htmlquery.InnerText(v))
					break
				case tradingCodeDSE:
					s.TradingCode = strings.TrimSpace(htmlquery.InnerText(v))
					break
				case ltpDSE:
					s.LTP = toFloat64(htmlquery.InnerText(v))
					break
				case highDSE:
					s.High = toFloat64(htmlquery.InnerText(v))
					break
				case lowDSE:
					s.Low = toFloat64(htmlquery.InnerText(v))
					break
				case closePriceDSE:
					s.CloseP = toFloat64(htmlquery.InnerText(v))
					break
				case ycpDSE:
					s.YCP = toFloat64(htmlquery.InnerText(v))
					break
				case changeDSE:
					s.Change = toFloat64(htmlquery.InnerText(v))
					break
				case tradeDSE:
					s.Trade = toInt64(htmlquery.InnerText(v))
					break
				case valueDSE:
					s.ValueInMN = toFloat64(htmlquery.InnerText(v))
					break
				case volumeDSE:
					s.Volume = toInt64(htmlquery.InnerText(v))
					break
				}

			}

			latestShares = append(latestShares, s)
		}
	}
	return latestShares, nil
}

// GetLatestPricesByCategory returns the array of latest share prices of the input category or error in case of any error
// It takes a category name, by which field the array should be sorted ex: SortByTradingCode and sort order ex: ASC
// It will return an error for if user tries to sort with a non existing file in the DSEShare model or invalid category name or invalid sort order
func (d *DSE) GetLatestPricesByCategory(categoryName string, by sortBy, order sortOrder) ([]*DSEShare, error) {
	categoryNameCap := strings.ToUpper(categoryName)
	if !isValidCategoryName(categoryNameCap) {
		return nil, errInvalidGroupName
	}

	url := fmt.Sprintf("https://www.dsebd.org/latest_share_price_all_group.php?group=%s", categoryNameCap)
	arr, err := getDSELatestPrices(url)
	if err != nil {
		return nil, err
	}

	arr, err = sortDse(arr, by, order)

	return arr, err
}

// GetLatestPrices returns the array of latest share prices or error in case of any error
// It takes by which field the array should be sorted ex: SortByTradingCode and sort order ex: ASC
// It will return an error for if user tries to sort with a non existing file in the DSEShare model or invalid category name or invalid sort order
func (d *DSE) GetLatestPrices(by sortBy, order sortOrder) ([]*DSEShare, error) {
	arr, err := getDSELatestPrices("https://www.dsebd.org/latest_share_price_scroll_l.php")
	if err != nil {
		return nil, err
	}
	arr, err = sortDse(arr, by, order)
	return arr, err
}

func sortDse(arr []*DSEShare, by sortBy, order sortOrder) ([]*DSEShare, error) {
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
	case SortByPriceChange:
		if order == ASC {
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].Change < arr[j].Change
			})
			return arr, nil
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Change > arr[j].Change
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

// LatestPricesWithPercentage ...
type LatestPricesWithPercentage struct {
	ID               int     `json:"id"`
	TradingCode      string  `json:"trading_code"`
	LTP              float64 `json:"ltp"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	CloseP           float64 `json:"close_p"`
	YCP              float64 `json:"ycp"`
	PercentageChange float64 `json:"percentage_change"`
	Trade            int64   `json:"trade"`
	ValueInMN        float64 `json:"value"`
	Volume           int64   `json:"volume"`
}

// GetLatestPricesSortedByPercentageChange ...
func (d *DSE) GetLatestPricesSortedByPercentageChange() ([]*LatestPricesWithPercentage, error) {
	latestPricesWithPercentage := make([]*LatestPricesWithPercentage, 0)
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/latest_share_price_all_by_change.php")
	if err != nil {
		return nil, err
	}

	tr, err := htmlquery.QueryAll(doc, "//tr")
	if err != nil {
		return nil, err
	}
	for i, t := range tr {
		if i == 0 {
			continue
		}
		td, err := htmlquery.QueryAll(t, "//td")
		if err != nil {
			return nil, err
		}
		s := &LatestPricesWithPercentage{}
		for index, v := range td {
			switch index {
			case 0:
				s.ID = toInt(htmlquery.InnerText(v))
				break
			case 1:
				s.TradingCode = strings.TrimSpace(htmlquery.InnerText(v))
				break
			case 2:
				s.LTP = toFloat64(htmlquery.InnerText(v))
				break
			case 3:
				s.High = toFloat64(htmlquery.InnerText(v))
				break
			case 4:
				s.Low = toFloat64(htmlquery.InnerText(v))
				break
			case 5:
				s.CloseP = toFloat64(htmlquery.InnerText(v))
				break
			case 6:
				s.YCP = toFloat64(htmlquery.InnerText(v))
				break
			case 7:
				s.PercentageChange = toFloat64(htmlquery.InnerText(v))
				break
			case 8:
				s.Trade = toInt64(htmlquery.InnerText(v))
				break
			case 9:
				s.ValueInMN = toFloat64(htmlquery.InnerText(v))
				break
			case 10:
				s.Volume = toInt64(htmlquery.InnerText(v))
				break
			}
		}
		latestPricesWithPercentage = append(latestPricesWithPercentage, s)
	}
	return latestPricesWithPercentage, nil
}

// DseMarketStatus holds the data for if market is open/close and when was last updated
type DseMarketStatus struct {
	IsOpen        bool
	LastUpdatedOn struct {
		Date string
		Time string
	}
}

// GetMarketStatus returns the DseMarketStatus with is open/close and last market update date time
func (d *DSE) GetMarketStatus() (*DseMarketStatus, error) {
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/")
	if err != nil {
		return nil, err
	}
	isOpenNode, err := htmlquery.Query(doc, `/html/body/div/div/div/header/div[1]/span[3]/span/b`)
	if err != nil {
		return nil, err
	}

	isOpenText := htmlquery.InnerText(isOpenNode)

	isOpen := false

	if isOpenText == "Open" {
		isOpen = true
	}

	dateTimeNode, err := htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/h2`)
	if err != nil {
		return nil, err
	}

	dateTimeText := htmlquery.InnerText(dateTimeNode)

	splitDateTime := strings.Split(dateTimeText, "Last update on ")[1]
	dateTime := strings.Split(splitDateTime, " at ")
	date := dateTime[0]
	time := dateTime[1]

	dseMarketStatus := &DseMarketStatus{
		IsOpen: isOpen,
		LastUpdatedOn: struct {
			Date string
			Time string
		}{Date: date, Time: time},
	}

	return dseMarketStatus, nil
}

// MarketSummary holds the data for market summary like DSEX index details and total trades and so on
type MarketSummary struct {
	LastUpdatedOn struct {
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"last_updated_on"`

	DseX struct {
		DSEXIndex                 float64 `json:"dsex_index"`
		DSEXIndexChange           float64 `json:"dsex_index_change"`
		DSEXIndexChangePercentage float64 `json:"dsex_index_change_percentage"`
	} `json:"dsex"`

	Ds30 struct {
		DS30Index                 float64 `json:"ds30_index"`
		DS30IndexChange           float64 `json:"ds30_index_change"`
		DS30IndexChangePercentage float64 `json:"ds30_index_change_percentage"`
	} `json:"ds30"`

	DseS struct {
		DSESIndex                 float64 `json:"dses_index"`
		DSESIndexChange           float64 `json:"dses_index_change"`
		DSESIndexChangePercentage float64 `json:"dses_index_change_percentage"`
	} `json:"dses"`

	TotalTrade     int64 `json:""`
	TotalValueInMN float64 `json:""`
	TotalVolume    int64   `json:""`

	IssuesAdvanced  int32 `json:""`
	IssuesDeclined  int32 `json:""`
	IssuesUnchanged int32 `json:""`
}

// GetMarketSummary returns the last updated market summary data
func (d *DSE) GetMarketSummary() (*MarketSummary, error) {
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/")
	if err != nil {
		return nil, err
	}

	dateTimeNode, err := htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/h2`)
	if err != nil {
		return nil, err
	}

	dateTimeText := htmlquery.InnerText(dateTimeNode)

	splitDateTime := strings.Split(dateTimeText, "Last update on ")[1]
	dateTime := strings.Split(splitDateTime, " at ")
	date := dateTime[0]
	time := dateTime[1]

	dseMarketSummary := &MarketSummary{}
	dseMarketSummary.LastUpdatedOn.Date = date
	dseMarketSummary.LastUpdatedOn.Time = time

	//var aggErr error
	node, err := htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/div[1]`)
	if err != nil {
		return nil, err
	}
	divNode, err := htmlquery.QueryAll(node, "//div")
	if err != nil {
		return nil, err
	}

	for index, v := range divNode {
		switch index {
		case 2:
			dsexIndexString := htmlquery.InnerText(v)
			dsexIndex, err := strconv.ParseFloat(strings.TrimSpace(dsexIndexString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseX.DSEXIndex = dsexIndex
		case 3:
			dsexIndexChangeString := htmlquery.InnerText(v)
			dsexIndexChange, err := strconv.ParseFloat(strings.TrimSpace(dsexIndexChangeString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseX.DSEXIndexChange = dsexIndexChange
		case 4:
			dsexIndexChangePString := htmlquery.InnerText(v)
			dsexIndexChangeP, err := strconv.ParseFloat(strings.TrimSpace(strings.Replace(dsexIndexChangePString,"%","",-1)) , 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseX.DSEXIndexChangePercentage = dsexIndexChangeP
		}
	}

	// DSES
	node, err = htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/div[2]`)
	if err != nil {
		return nil, err
	}
	divNode, err = htmlquery.QueryAll(node, "//div")
	if err != nil {
		return nil, err
	}

	for index, v := range divNode {
		switch index {
		case 2:
			dsesIndexString := htmlquery.InnerText(v)
			dsesIndex, err := strconv.ParseFloat(strings.TrimSpace(dsesIndexString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseS.DSESIndex = dsesIndex
		case 3:
			dsesIndexChangeString := htmlquery.InnerText(v)
			dsesIndexChange, err := strconv.ParseFloat(strings.TrimSpace(dsesIndexChangeString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseS.DSESIndexChange = dsesIndexChange
		case 4:
			dsesIndexChangePString := htmlquery.InnerText(v)
			dsesIndexChangeP, err := strconv.ParseFloat(strings.TrimSpace(strings.Replace(dsesIndexChangePString,"%","",-1)) , 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.DseS.DSESIndexChangePercentage = dsesIndexChangeP
		}
	}


	// DS30
	node, err = htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/div[3]`)
	if err != nil {
		return nil, err
	}
	divNode, err = htmlquery.QueryAll(node, "//div")
	if err != nil {
		return nil, err
	}

	for index, v := range divNode {
		switch index {
		case 2:
			ds30IndexString := htmlquery.InnerText(v)
			ds30Index, err := strconv.ParseFloat(strings.TrimSpace(ds30IndexString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.Ds30.DS30Index = ds30Index
		case 3:
			ds30IndexChangeString := htmlquery.InnerText(v)
			ds30IndexChange, err := strconv.ParseFloat(strings.TrimSpace(ds30IndexChangeString), 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.Ds30.DS30IndexChange = ds30IndexChange
		case 4:
			ds30IndexChangePString := htmlquery.InnerText(v)
			ds30IndexChangeP, err := strconv.ParseFloat(strings.TrimSpace(strings.Replace(ds30IndexChangePString,"%","",-1)) , 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.Ds30.DS30IndexChangePercentage = ds30IndexChangeP
		}
	}

	// Total Trade, Volumes, Value
	node, err = htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/div[5]`)
	if err != nil {
		return nil, err
	}
	divNode, err = htmlquery.QueryAll(node, "//div")
	if err != nil {
		return nil, err
	}

	for index, v := range divNode {
		switch index {
		case 1:
			totalTradeString := htmlquery.InnerText(v)
			totalTrade, err := strconv.ParseInt(strings.TrimSpace(totalTradeString), 10, 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.TotalTrade = totalTrade
		case 2:
			totalVolumeString := htmlquery.InnerText(v)
			totalVolume, err := strconv.ParseInt(strings.TrimSpace(totalVolumeString), 10, 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.TotalVolume = totalVolume
		case 3:
			totalValueString := htmlquery.InnerText(v)
			totalValue, err := strconv.ParseFloat(strings.TrimSpace(strings.Replace(totalValueString,"%","",-1)) , 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.Ds30.DS30IndexChangePercentage = totalValue
		}
	}

	// Total issues
	node, err = htmlquery.Query(doc, `/html/body/div[2]/section/div/div[1]/div/div[7]`)
	if err != nil {
		return nil, err
	}
	divNode, err = htmlquery.QueryAll(node, "//div")
	if err != nil {
		return nil, err
	}

	for index, v := range divNode {
		switch index {
		case 1:
			issuesAdvancedString := htmlquery.InnerText(v)
			issuesAdvanced, err := strconv.ParseInt(strings.TrimSpace(issuesAdvancedString), 10, 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.IssuesAdvanced = int32(issuesAdvanced)
		case 2:
			issuesDeclinedString := htmlquery.InnerText(v)
			issuesDeclined, err := strconv.ParseInt(strings.TrimSpace(issuesDeclinedString), 10, 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.IssuesDeclined = int32(issuesDeclined)
		case 3:
			issuesUnchangedString := htmlquery.InnerText(v)
			issuesUnchanged, err := strconv.ParseInt(strings.TrimSpace(strings.Replace(issuesUnchangedString,"%","",-1)), 10, 64)
			if err != nil {
				return nil, err
			}
			dseMarketSummary.IssuesUnchanged = int32(issuesUnchanged)
		}
	}

	return dseMarketSummary, nil
}
