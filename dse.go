package bdstockexchange

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"sort"
	"strings"
)

type dse struct {
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

func NewDSE() *dse {
	return new(dse)
}

type dseShare struct {
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

func getDSELatestPrices(url string) ([]*dseShare, error) {
	latestShares := make([]*dseShare, 0)
	// Request the HTML page.
	if url == "" {
		url = "https://www.dsebd.org/latest_share_price_all.php"
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	doc.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		//log.Println(selection.Find("td").Text())
		if i == 0 {
			return
		}
		s := &dseShare{}
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

func (d *dse) GetLatestPricesByCategory(categoryName string) ([]*dseShare, error) {
	categoryNameCap := strings.ToUpper(categoryName)
	if !isValidCategoryName(categoryNameCap) {
		return nil, ErrInvalidGroupName
	}

	url := fmt.Sprintf("https://www.dsebd.org/latest_share_price_all_group.php?group=%s", categoryNameCap)
	arr, err := getDSELatestPrices(url)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (d *dse) GetLatestPricesSortedByTradingCode() ([]*dseShare, error) {
	arr, err := getDSELatestPrices("https://www.dsebd.org/latest_share_price_all.php")
	if err != nil {
		return nil, err
	}
	return arr, err
}

func (d *dse) GetLatestPricesSortedByChange() ([]*dseShare, error) {
	arr, err := getDSELatestPrices("")
	if err != nil {
		return nil, err
	}
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i].Change > arr[j].Change
	})
	return arr, nil
}
