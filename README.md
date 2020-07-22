## bdstockexchange
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

#### type CSEShare

```go
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
```

CSEShare is a model for a single company's latest price data provided by the cse
website

#### func (*CSEShare) GetLatestPrices

```go
func (c *CSEShare) GetLatestPrices(by sortBy, order sortOrder) ([]*CSEShare, error)
```
GetLatestPrices returns the array of latest share prices or error in case of any
error It takes by which field the array should be sorted ex: SortByTradingCode
and sort order ex: ASC It will return an error for if user tries to sort with a
non existing file in the CSEShare model or invalid category name or invalid sort
order

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
