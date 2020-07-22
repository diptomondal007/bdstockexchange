package bdstockexchange

// sortBy is the type for constants that can be used to pass as a parameter type of sort wanted by the user
type sortBy uint8

// sortOrder ...
type sortOrder uint8

const (
	// ASC constant to sort result array in ascending order
	ASC sortOrder = iota
	// DESC constant to sort result array in descending order
	DESC
)

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
