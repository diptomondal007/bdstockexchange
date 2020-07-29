## bdstockexchange
A Go library to fetch latest stock prices from Dhaka and Chittagong Stock Exchange (DSE & CSE).

[![Build Status](https://travis-ci.com/diptomondal007/bdstockexchange.svg?branch=master)](https://travis-ci.com/github/diptomondal007/bdstockexchange)
[![Coverage Status](https://coveralls.io/repos/github/diptomondal007/bdstockexchange/badge.svg?branch=master)](https://coveralls.io/github/diptomondal007/bdstockexchange?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/diptomondal007/bdstockexchange)](https://goreportcard.com/report/github.com/diptomondal007/bdstockexchange)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/diptomondal007/bdstockexchange?tab=doc)](https://pkg.go.dev/github.com/diptomondal007/bdstockexchange?tab=doc)

## Import
```go
import "github.com/diptomondal007/bdstockexchange"
```
## Install
```
go get -u github.com/diptomondal007/bdstockexchange
```

## Usage

```go
const (
	// ASC constant to sort result array in ascending order
	ASC sortOrder = iota
	// DESC constant to sort result array in descending order
	DESC
)
```

```go
const (
	// SortByTradingCode to sort the result by Company's Trade code
	SortByTradingCode sortBy = iota
	// SortByLTP to sort the result by the Last Trade Price
	SortByLTP
	// SortByOpeningPrice to sort the result by the Opening Price of that day
	SortByOpeningPrice
	// SortByHighPrice to sort the result by the Highest Price of the day
	SortByHighPrice
	// SortByLowPrice to sort the result by Lowest price of the day
	SortByLowPrice
	// SortByYCP to sort the result by Yesterday's Closing Price
	SortByYCP
	// SortByNumberOfTrades to sort the result by The Number of shares are traded on that day
	SortByNumberOfTrades
	// SortByValue to sort the result by the Value of the Company. The Value is in Million BDT.
	SortByValue
	// SortByVolumeOfShare to sort the result by the Number of shares of the company
	SortByVolumeOfShare
	// SortByPercentageChange ...
	SortByPercentageChange
	// SortByPriceChange to sort the result by the Change of Price of the Share
	SortByPriceChange
)
```

#### type CSE

```go
type CSE struct {
}
```

CSE is a struct to access cse related methods

#### func  NewCSE

```go
func NewCSE() *CSE
```
NewCSE returns new CSE object

#### func (*CSE) GetAllListedCompanies

```go
func (c *CSE) GetAllListedCompanies() ([]*Company, error)
```
GetAllListedCompanies returns all the companies listed in cse or error in case
of any error

#### func (*CSE) GetAllListedCompaniesByCategory

```go
func (c *CSE) GetAllListedCompaniesByCategory() ([]*CompanyListingByCategory, error)
```
GetAllListedCompaniesByCategory returns the listing of the companies by their
category or an error in case of any error

#### func (*CSE) GetAllListedCompaniesByIndustry

```go
func (c *CSE) GetAllListedCompaniesByIndustry() ([]*CompanyListingByIndustry, error)
```
GetAllListedCompaniesByIndustry returns list of companies with their industry
type or error in case of any error

#### func (*CSE) GetAllWeeklyReports

```go
func (c *CSE) GetAllWeeklyReports(year int) (*WeeklyReports, error)
```
GetAllWeeklyReports returns weekly reports pdf link for the input Year. the Year
should be between current Year and 2018

#### func (*CSE) GetLatestPrices

```go
func (c *CSE) GetLatestPrices(by sortBy, order sortOrder) ([]*CSEShare, error)
```
GetLatestPrices returns the array of latest share prices or error in case of any
error It takes by which field the array should be sorted ex: SortByTradingCode
and sort order ex: ASC It will return an error for if user tries to sort with a
non existing file in the CSEShare model or invalid category name or invalid sort
order

#### func (*CSE) GetMarketStatus

```go
func (d *CSE) GetMarketStatus() (*CseMarketStatus, error)
```
GetMarketStatus returns the CseMarketStatus with is open/close

#### func (*CSE) GetMarketSummary

```go
func (c *CSE) GetMarketSummary() (*Summary, error)
```
GetMarketSummary returns the summary with highest records till now and the
historical market summary data

#### func (*CSE) GetPriceEarningRatio

```go
func (c *CSE) GetPriceEarningRatio(day, month, year string) (*PriceEarningRatios, error)
```
GetPriceEarningRatio returns the price earning ratio data for listed companies
as per input date. It takes day, month and Year as input ex : (03, 07, 2020)
where 03 is the day and 07 is the month and 2020 is the Year. Don't forget to
include 0 before single digit day or month

#### type CSEShare

```go
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
```

CSEShare is a model for a single company's latest price data provided by the cse
website

#### type Company

```go
type Company struct {
	CompanyName string
	TradingCode string
}
```


#### type CompanyListingByCategory

```go
type CompanyListingByCategory struct {
	Category string
	List     []*Company
}
```


#### type CompanyListingByIndustry

```go
type CompanyListingByIndustry struct {
	IndustryType string
	List         []*Company
}
```


#### type CseMarketStatus

```go
type CseMarketStatus struct {
	IsOpen bool
}
```

CseMarketStatus holds the data for if market is open/close

#### type DSE

```go
type DSE struct {
}
```

DSE is a struct to access dse related methods

#### func  NewDSE

```go
func NewDSE() *DSE
```
NewDSE returns new DSE object

#### func (*DSE) GetLatestPrices

```go
func (d *DSE) GetLatestPrices(by sortBy, order sortOrder) ([]*DSEShare, error)
```
GetLatestPrices returns the array of latest share prices or error in case of any
error It takes by which field the array should be sorted ex: SortByTradingCode
and sort order ex: ASC It will return an error for if user tries to sort with a
non existing file in the DSEShare model or invalid category name or invalid sort
order

#### func (*DSE) GetLatestPricesByCategory

```go
func (d *DSE) GetLatestPricesByCategory(categoryName string, by sortBy, order sortOrder) ([]*DSEShare, error)
```
GetLatestPricesByCategory returns the array of latest share prices of the input
category or error in case of any error It takes a category name, by which field
the array should be sorted ex: SortByTradingCode and sort order ex: ASC It will
return an error for if user tries to sort with a non existing file in the
DSEShare model or invalid category name or invalid sort order

#### func (*DSE) GetLatestPricesSortedByPercentageChange

```go
func (d *DSE) GetLatestPricesSortedByPercentageChange() ([]*LatestPricesWithPercentage, error)
```
GetLatestPricesSortedByPercentageChange ...

#### func (*DSE) GetMarketStatus

```go
func (d *DSE) GetMarketStatus() (*DseMarketStatus, error)
```
GetMarketStatus returns the DseMarketStatus with is open/close and last market
update date time

#### type DSEShare

```go
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
```

DSEShare is a model for a single company's latest price data provided by the dse
website

#### type DseMarketStatus

```go
type DseMarketStatus struct {
	IsOpen        bool
	LastUpdatedOn struct {
		Date string
		Time string
	}
}
```

DseMarketStatus holds the data for if market is open/close and when was last
updated

#### type LatestPricesWithPercentage

```go
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
```

LatestPricesWithPercentage ...

#### type PriceEarningRatio

```go
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
```

PriceEarningRatio holds the data for a price earning ratio in selected date

#### type PriceEarningRatios

```go
type PriceEarningRatios struct {
	Date                   string
	PriceEarningRatioArray []*PriceEarningRatio
}
```


#### type Summary

```go
type Summary struct {
	HighestRecords      []*record
	HistoricalSummaries []*market
}
```

Summary holds the historical market summaries array and the record trading or
highest records data

#### type WeeklyReports

```go
type WeeklyReports struct {
	Year    int
	Reports []*report
}
```

WeeklyReports holds the weekly reports for a Year


## Example
#### GetLatestPrices
```go
package main

import (
	"github.com/diptomondal007/bdstockexchange"
	"log"
)

func main(){
	dse := bdstockexchange.NewDSE()
	arr, err := dse.GetLatestPrices(bdstockexchange.SortByHighPrice, bdstockexchange.ASC)
	if err != nil{
		// Do something with the error
		log.Println(err)
	}
	log.Println(arr[0].TradingCode)
}
```

#### GetLatestPricesByCategory
```go
package main

import (
	"github.com/diptomondal007/bdstockexchange"
	"log"
)

func main(){
	dse := bdstockexchange.NewDSE()
	arr, err := dse.GetLatestPricesByCategory("A" ,bdstockexchange.SortByHighPrice, bdstockexchange.ASC)
	if err != nil{
		// Do something with the error
		log.Println(err)
	}
	log.Println(arr[0].TradingCode)
}
```

## License

bdstockexchange is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/diptomondal007/bdstockexchange/blob/master/LICENSE)
