package bdstockexchange

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		url = "https://www.dsebd.org/latest_share_price_all.php"
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		//log.Println(selection.Find("td").Text())
		if i == 0 {
			return
		}
		s := &DSEShare{}
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case idDSE:
				s.ID = toInt(selection.Text())
				break
			case tradingCodeDSE:
				s.TradingCode = strings.TrimSpace(selection.Text())
				break
			case ltpDSE:
				s.LTP = toFloat64(selection.Text())
				break
			case highDSE:
				s.High = toFloat64(selection.Text())
				break
			case lowDSE:
				s.Low = toFloat64(selection.Text())
				break
			case closePriceDSE:
				s.CloseP = toFloat64(selection.Text())
				break
			case ycpDSE:
				s.YCP = toFloat64(selection.Text())
				break
			case changeDSE:
				s.Change = toFloat64(selection.Text())
				break
			case tradeDSE:
				s.Trade = toInt64(selection.Text())
				break
			case valueDSE:
				s.ValueInMN = toFloat64(selection.Text())
				break
			case volumeDSE:
				s.Volume = toInt64(selection.Text())
				break
			}
		})
		latestShares = append(latestShares, s)
	})
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
	arr, err := getDSELatestPrices("https://www.dsebd.org/latest_share_price_all.php")
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

	dateTimeNode, err := htmlquery.Query(doc, `/html/body/div/div/div/div[1]/div[1]/h2`)
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
