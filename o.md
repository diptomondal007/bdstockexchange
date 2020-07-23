# bdstockexchange
--
    import "."


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

#### func (*CSE) GetLatestPrices

```go
func (c *CSE) GetLatestPrices(by sortBy, order sortOrder) ([]*CSEShare, error)
```
GetLatestPrices returns the array of latest share prices or error in case of any
error It takes by which field the array should be sorted ex: SortByTradingCode
and sort order ex: ASC It will return an error for if user tries to sort with a
non existing file in the CSEShare model or invalid category name or invalid sort
order



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


