package bdstockexchange

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"

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

// WeeklyReports holds the weekly reports for a Year
type WeeklyReports struct {
	Year    int
	Reports []*report
}

type report struct {
	Date          string
	Title         string
	ReportPDFLink string
}

type Company struct {
	CompanyName string
	TradingCode string
}

type CompanyListingByIndustry struct {
	IndustryType string
	List         []*Company
}

type CompanyListingByCategory struct {
	Category string
	List     []*Company
}

type PriceEarningRatios struct {
	Date                   string
	PriceEarningRatioArray []*PriceEarningRatio
}

// PriceEarningRatio holds the data for a price earning ratio in selected date
type PriceEarningRatio struct {
	SL            string
	TradingCode   string
	FinancialYear struct {
		From string
		To   string
	}
	EPSAsPerUpdatedUnAuditedAccounts struct {
		Quarter1 float64
		HalfYear float64
		Quarter3 float64
	}
	AnnualizedEPS                     float64
	EPSBasedOnLastAuditedAccounts     float64
	ClosePrice                        float64
	PERatioBasedOnAnnualizedEPS       float64
	PERatioBasedOnLastAuditedAccounts float64
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

//NewCSE returns new CSE object
func NewCSE() *CSE {
	return new(CSE)
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

// GetAllWeeklyReports returns weekly reports pdf link for the input Year. the Year should be between current Year and 2018
func (c *CSE) GetAllWeeklyReports(year int) (*WeeklyReports, error) {
	data := fmt.Sprintf("Year=%d", year)
	body := strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://www.cse.com.bd/market/weekly_report", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "www.cse.com.bd")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://www.cse.com.bd")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Mobile Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://www.cse.com.bd/market/weekly_report")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "logins=281d8de2fe59ebd74a2fb76b0f75bcbf229bd7f1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	//defer resp.Body.Close()

	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))

	if err != nil {
		// handle err
	}

	doc, err := html.Parse(r)

	if err != nil {
		// handle error
	}

	reports := make([]*report, 0)
	weeklyReports := &WeeklyReports{
		Year:    year,
		Reports: nil,
	}

	availableYears := make([]int, 0)
	list := htmlquery.Find(doc, `//*[@id="wrapper"]/div/div/div[1]/div/div[1]/div/div/form/div/div[2]/select`)
	for _, v := range list {
		option := htmlquery.Find(v, "option")
		for _, v := range option {
			if htmlquery.InnerText(v) == "" {
				continue
			} else {
				availableYears = append(availableYears, toInt(htmlquery.InnerText(v)))
			}
		}
	}

	isValidYear := false

	for _, v := range availableYears {
		if v == year {
			isValidYear = true
			break
		}
	}

	if isValidYear == false {
		return nil, errNotAValidYear
	}

	list = htmlquery.Find(doc, `//*[@id="wrapper"]/div/div/div[1]/div/div[3]/div/div/div[2]/div`)

	for _, v := range list {
		peRatioTabsContent := htmlquery.Find(v, "div")
		for _, v := range peRatioTabsContent {
			if htmlquery.SelectAttr(v, "class") == "pe_ratio_tabs_cont" {
				date := htmlquery.InnerText(htmlquery.FindOne(v, `//*[@id="pe_ratiocont_1"]`))
				titleNode := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_2"]`)
				title := htmlquery.InnerText(titleNode)
				link := htmlquery.InnerText(htmlquery.FindOne(titleNode, "//a/@href"))

				rep := &report{
					Date:          date,
					Title:         title,
					ReportPDFLink: link,
				}
				reports = append(reports, rep)
			}
		}
	}

	weeklyReports.Reports = reports

	return weeklyReports, nil
}

// GetAllListedCompanies returns all the companies listed in cse or error in case of any error
func (c *CSE) GetAllListedCompanies() ([]*Company, error) {
	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/company/listedcompanies")
	if err != nil {
		return nil, err
	}

	companyListing := make([]*Company, 0)

	list, err := htmlquery.QueryAll(doc, `//*[@id="top_content_1"]/div/div/div/div/div/div/div[2]`)
	if err != nil {
		return nil, err
	}
	for _, l := range list {
		div := htmlquery.Find(l, "//div//ul")
		for _, di := range div {
			li := htmlquery.Find(di, "li")
			for _, v := range li {
				a := htmlquery.FindOne(v, "a")
				companyName := strings.TrimSpace(htmlquery.InnerText(a))
				companyTradingCode := strings.Trim(htmlquery.InnerText(htmlquery.FindOne(v, "//a/@href")), "https://www.cse.com.bd/company/companydetails/")

				company := &Company{
					CompanyName: companyName,
					TradingCode: companyTradingCode,
				}
				companyListing = append(companyListing, company)
			}

		}
	}

	return companyListing, nil
}

// GetAllListedCompaniesByIndustry returns list of companies with their industry type or error in case of any error
func (c *CSE) GetAllListedCompaniesByIndustry() ([]*CompanyListingByIndustry, error) {
	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/company/listedcompanies")
	if err != nil {
		return nil, err
	}

	companyListIndustry := make([]*CompanyListingByIndustry, 0)

	list, err := htmlquery.QueryAll(doc, `//*[@id="top_content_2"]/div/div/div/div/div`)
	if err != nil {
		return nil, err
	}

	for _, li := range list {
		divs, err := htmlquery.QueryAll(li, "div")
		//log.Println(htmlquery.InnerText(htmlquery.FindOne(divs, "//@id")))
		if err != nil {
			return nil, err
		}
		for _, div := range divs {
			list := make([]*Company, 0)
			category := htmlquery.FindOne(div, "//@id")

			li := htmlquery.Find(div, "//ul//li")
			for _, v := range li {
				a := htmlquery.FindOne(v, "a")
				companyName := strings.TrimSpace(htmlquery.InnerText(a))
				companyTradingCode := strings.Trim(htmlquery.InnerText(htmlquery.FindOne(v, "//a/@href")), "https://www.cse.com.bd/company/companydetails/")

				company := &Company{
					CompanyName: companyName,
					TradingCode: companyTradingCode,
				}
				list = append(list, company)
			}
			listByIndustry := &CompanyListingByIndustry{
				IndustryType: htmlquery.InnerText(category),
				List:         list,
			}
			companyListIndustry = append(companyListIndustry, listByIndustry)
		}
	}
	return companyListIndustry, nil
}

// GetAllListedCompaniesByCategory returns the listing of the companies by their category or an error in case of any error
func (c *CSE) GetAllListedCompaniesByCategory() ([]*CompanyListingByCategory, error) {
	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/company/listedcompanies")
	if err != nil {
		return nil, err
	}

	companyListByCategory := make([]*CompanyListingByCategory, 0)

	list, err := htmlquery.QueryAll(doc, `//*[@id="top_content_3"]/div/div/div/div/div`)

	if err != nil {
		return nil, err
	}

	for _, li := range list {
		divs, err := htmlquery.QueryAll(li, "div")
		//log.Println(htmlquery.InnerText(htmlquery.FindOne(divs, "//@id")))
		if err != nil {
			return nil, err
		}

		for _, div := range divs {
			list := make([]*Company, 0)
			category := htmlquery.FindOne(div, "div")

			li := htmlquery.Find(div, "//ul//li")
			for _, v := range li {
				a := htmlquery.FindOne(v, "a")
				companyName := strings.TrimSpace(htmlquery.InnerText(a))
				companyTradingCode := strings.Trim(htmlquery.InnerText(htmlquery.FindOne(v, "//a/@href")), "https://www.cse.com.bd/company/companydetails/")

				company := &Company{
					CompanyName: companyName,
					TradingCode: companyTradingCode,
				}
				list = append(list, company)
			}
			listByCategory := &CompanyListingByCategory{
				Category: strings.TrimSpace(htmlquery.InnerText(category)),
				List:     list,
			}
			companyListByCategory = append(companyListByCategory, listByCategory)
		}
	}
	return companyListByCategory, nil
}

// GetPriceEarningRatio returns the price earning ratio data for listed companies as per input date. It takes day, month and Year as input ex : (03, 07, 2020)
// where 03 is the day and 07 is the month and 2020 is the Year. Don't forget to include 0 before single digit day or month
func (c *CSE) GetPriceEarningRatio(day, month, year string) (*PriceEarningRatios, error) {
	priceEarningRatios := &PriceEarningRatios{
		Date:                   "",
		PriceEarningRatioArray: nil,
	}

	priceEarningRatioArray := make([]*PriceEarningRatio, 0)

	data := fmt.Sprintf("pe_date=%s-%s-%s", year, month, day)
	data_body := strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://www.cse.com.bd/market/pe_ratio", data_body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authority", "www.cse.com.bd")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://www.cse.com.bd")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Mobile Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://www.cse.com.bd/market/pe_ratio")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "logins=e879d1d3a477b3e43ddbb30e4d46ad4feddddfea")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//re, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(re))

	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	list, err := htmlquery.QueryAll(doc, `//*[@id="wrapper"]/div/div/div[1]/div/div[3]/div/div/div/div`)
	if err != nil {
		return nil, err
	}

	var isDataFound bool

	for _, v := range list {
		tabsContents, err := htmlquery.QueryAll(v, "//div")
		if err != nil {
			return nil, err
		}
		isDataFound = false
		for _, v := range tabsContents {
			if htmlquery.SelectAttr(v, "class") == "pe_ratio_tabs_cont" {
				isDataFound = true
				pe_ratiocont_1 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_1"]`)
				pe_ratiocont_2 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_2"]`)
				pe_ratiocont_3_td1 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_3"]/table/tbody/tr/td[1]`)
				pe_ratiocont_3_td2 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_3"]/table/tbody/tr/td[2]`)
				//log.Println(htmlquery.InnerText(pe_ratiocont_3_td1), htmlquery.InnerText(pe_ratiocont_3_td2))
				pe_ratiocont_4_td1 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_4"]/table/tbody/tr/td[1]`)
				pe_ratiocont_4_td2 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_4"]/table/tbody/tr/td[2]`)
				pe_ratiocont_4_td3 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_4"]/table/tbody/tr/td[3]`)
				pe_ratiocont_5 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_5"]`)
				pe_ratiocont_6 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_6"]`)
				pe_ratiocont_7 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_7"]`)
				pe_ratiocont_8 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_8"]`)
				pe_ratiocont_9 := htmlquery.FindOne(v, `//*[@id="pe_ratiocont_9"]`)

				priceEarningRatio := &PriceEarningRatio{
					SL:          strings.Replace(htmlquery.InnerText(pe_ratiocont_1), ".", "", -1),
					TradingCode: htmlquery.InnerText(pe_ratiocont_2),
					FinancialYear: struct {
						From string
						To   string
					}{
						From: htmlquery.InnerText(pe_ratiocont_3_td1),
						To:   htmlquery.InnerText(pe_ratiocont_3_td2),
					},
					EPSAsPerUpdatedUnAuditedAccounts: struct {
						Quarter1 float64
						HalfYear float64
						Quarter3 float64
					}{
						Quarter1: toFloat64(htmlquery.InnerText(pe_ratiocont_4_td1)),
						HalfYear: toFloat64(htmlquery.InnerText(pe_ratiocont_4_td2)),
						Quarter3: toFloat64(htmlquery.InnerText(pe_ratiocont_4_td3)),
					},
					AnnualizedEPS:                     toFloat64(htmlquery.InnerText(pe_ratiocont_5)),
					EPSBasedOnLastAuditedAccounts:     toFloat64(htmlquery.InnerText(pe_ratiocont_6)),
					ClosePrice:                        toFloat64(htmlquery.InnerText(pe_ratiocont_7)),
					PERatioBasedOnAnnualizedEPS:       toFloat64(htmlquery.InnerText(pe_ratiocont_8)),
					PERatioBasedOnLastAuditedAccounts: toFloat64(htmlquery.InnerText(pe_ratiocont_9)),
				}

				priceEarningRatioArray = append(priceEarningRatioArray, priceEarningRatio)

			}
		}

	}
	if !isDataFound{
		return nil, errNoDataFound
	}

	priceEarningRatios.Date = data
	priceEarningRatios.PriceEarningRatioArray = priceEarningRatioArray

	return priceEarningRatios, nil
}

// CseMarketStatus holds the data for if market is open/close
type CseMarketStatus struct {
	IsOpen        bool
}

// GetMarketStatus returns the CseMarketStatus with is open/close
func (d *CSE) GetMarketStatus() (*CseMarketStatus, error) {
	doc, err := htmlquery.LoadURL("https://www.cse.com.bd/market/current_price")
	if err != nil {
		return nil, err
	}
	isOpenNode, err := htmlquery.Query(doc, `//*[@id="wrapper"]/div/header/div/div/div[2]/div[1]/div[1]/span`)
	if err != nil {
		return nil, err
	}

	isOpenText := htmlquery.InnerText(isOpenNode)

	isOpen := false

	if isOpenText == "Open" {
		isOpen = true
	}

	cseMarketStatus := &CseMarketStatus{
		IsOpen: isOpen,
	}

	log.Println(cseMarketStatus.IsOpen)

	return cseMarketStatus, nil
}
