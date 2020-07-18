package bd_sotck_exchange

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	colly "github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
)

const (
	ID = iota
	TradingCode
	LTP
	High
	Low
	CloseP
	YCP
	Change
	Trade
	Value
	Volume
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

func (d *dse) GetLatestPricesByTradeCode() ([]*latest, error) {
	latestShares := make([]*latest, 0)
	doc, err := htmlquery.LoadURL("https://www.dsebd.org/latest_share_price_all.php")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	list, err := htmlquery.QueryAll(doc, "//table //tr")
	if err != nil {
		panic(err)
		return nil, err
	}
	//log.Println(len(list))
	for i, n := range list {
		//log.Println(a)
		if i == 0 {
			continue
		}
		a := htmlquery.Find(n, "//td")
		s := &latest{
			ID:          htmlquery.InnerText(a[ID]),
			TradingCode: strings.TrimSpace(htmlquery.InnerText(a[TradingCode])),
			LTP:         htmlquery.InnerText(a[LTP]),
			High:        htmlquery.InnerText(a[High]),
			Low:         htmlquery.InnerText(a[Low]),
			CloseP:      htmlquery.InnerText(a[CloseP]),
			YCP:         htmlquery.InnerText(a[YCP]),
			Change:      htmlquery.InnerText(a[Change]),
			Trade:       htmlquery.InnerText(a[Trade]),
			Value:       htmlquery.InnerText(a[Value]),
			Volume:      htmlquery.InnerText(a[Volume]),
		}
		latestShares = append(latestShares, s)
	}
	return latestShares, nil
}

func (d *dse) GetLatestPricesByCategory(categoryName string) ([]*latest, error){
	categoryNameCap := strings.ToUpper(categoryName)
	if !isValidGroupName(categoryNameCap){
		return nil, ErrInvalidGroupName
	}
	url := fmt.Sprintf("https://www.dsebd.org/latest_share_price_all_group.php?group=%s", categoryNameCap)
	latestShares := make([]*latest, 0)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	list, err := htmlquery.QueryAll(doc, "//table //tr")
	if err != nil {
		return nil, err
	}
	//log.Println(len(list))

	for i, n := range list {
		//log.Println(a)
		if i == 0 {
			continue
		}
		a := htmlquery.Find(n, "//td")
		s := &latest{
			ID:          htmlquery.InnerText(a[ID]),
			TradingCode: strings.TrimSpace(htmlquery.InnerText(a[TradingCode])),
			LTP:         htmlquery.InnerText(a[LTP]),
			High:        htmlquery.InnerText(a[High]),
			Low:         htmlquery.InnerText(a[Low]),
			CloseP:      htmlquery.InnerText(a[CloseP]),
			YCP:         htmlquery.InnerText(a[YCP]),
			Change:      htmlquery.InnerText(a[Change]),
			Trade:       htmlquery.InnerText(a[Trade]),
			Value:       htmlquery.InnerText(a[Value]),
			Volume:      htmlquery.InnerText(a[Volume]),
		}
		latestShares = append(latestShares, s)
	}
	return latestShares, nil
}

func(d *dse) GetLatestMarketPriceByTradingCode(tradingCode string){

	req, err := http.NewRequest("POST", "https://www.dsebd.org/bshis_new1_nf.php", strings.NewReader("inst=AAMRANET"))
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Set("Origin", "https://www.dsebd.org")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.dsebd.org/mkt_depth_3.php")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "__gads=ID=ae0e2e143b4c279b:T=1594183361:S=ALNI_Mb9RAfjR4Ch8275xA84ULPfheb6Qg; _ga=GA1.2.2020771556.1594183361; _gid=GA1.2.1073117772.1594836039; _gat=1")
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	// handle err
	//}
	//defer resp.Body.Close()
	//r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	//if err != nil{
	//	log.Println(err)
	//}
	//parsedHTML, err := html.Parse(r)
	//list, err := htmlquery.QueryAll(parsedHTML, "//table //tr")
	//if err != nil {
	//	log.Println(err)
	//}
	////log.Println(len(list))
	//
	//for _, n := range list {
	//	//log.Println(n)
	//	a := htmlquery.Find(n, "//td")
	//	log.Println(htmlquery.InnerText(a[1]))
	//}

	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		//log.Println(string(response.Body))
	})
	c.OnHTML("tr", func(element *colly.HTMLElement) {
		element.ForEach("td", func(_ int, element *colly.HTMLElement) {
			log.Println(strings.TrimSpace(element.Text))
		})
	})
	err = c.Request("POST", "https://www.dsebd.org/bshis_new1_nf.php", strings.NewReader("inst=AAMRANET"), nil, req.Header)
}