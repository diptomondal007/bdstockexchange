package bdstockexchange

import (
	"reflect"
	"testing"
)

func TestNewCSE(t *testing.T) {
	tests := []struct {
		name string
		want *CSE
	}{
		// TODO: Add test cases.
		{"new",new(CSE)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSE(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCSELatestPrices(t *testing.T) {
	tests := []struct {
		name    string
		want    []*CSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", make([]*CSEShare,329), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCSELatestPrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCSELatestPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("getCSELatestPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSE_GetLatestPrices(t *testing.T) {
	type args struct {
		by    sortBy
		order sortOrder
	}
	tests := []struct {
		name    string
		c       *CSE
		args    args
		want    []*CSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", NewCSE() ,args{SortByHighPrice, ASC}, make([]*CSEShare,329), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSE{}
			got, err := c.GetLatestPrices(tt.args.by, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSE.GetLatestPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("CSE.GetLatestPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortCse(t *testing.T) {
	type args struct {
		arr   []*CSEShare
		by    sortBy
		order sortOrder
	}
	tests := []struct {
		name    string
		args    args
		want    []*CSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{make([]*CSEShare,0), SortByTradingCode, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByTradingCode, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByHighPrice, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByHighPrice, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByLowPrice, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByLowPrice, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByLTP, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByLTP, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByNumberOfTrades, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByNumberOfTrades, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByYCP, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByYCP, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByValue, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByValue, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByVolumeOfShare, ASC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByVolumeOfShare, DESC}, make([]*CSEShare,0), false},
		{"", args{make([]*CSEShare,0), SortByPercentageChange, DESC}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortCse(tt.args.arr, tt.args.by, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortCse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortCse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSE_GetMarketSummary(t *testing.T) {
	tests := []struct {
		name    string
		c       *CSE
		want    *Summary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSE{}
			got, err := c.GetMarketSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("CSE.GetMarketSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSE.GetMarketSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
