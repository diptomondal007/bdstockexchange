package bdstockexchange

import (
	"reflect"
	"testing"
)

func TestNewDSE(t *testing.T) {
	tests := []struct {
		name string
		want *DSE
	}{
		// TODO: Add test cases.
		{"new",new(DSE)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDSE(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDSELatestPrices(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []*DSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",args{""}, make([]*DSEShare,359), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDSELatestPrices(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDSELatestPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("getDSELatestPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSE_GetLatestPricesByCategory(t *testing.T) {
	type args struct {
		categoryName string
		by           sortBy
		order        sortOrder
	}
	tests := []struct {
		name    string
		d       *DSE
		args    args
		want    []*DSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", NewDSE() ,args{"A", SortByHighPrice, ASC}, make([]*DSEShare,258), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DSE{}
			got, err := d.GetLatestPricesByCategory(tt.args.categoryName, tt.args.by, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSE.GetLatestPricesByCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("DSE.GetLatestPricesByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSE_GetLatestPrices(t *testing.T) {
	type args struct {
		by    sortBy
		order sortOrder
	}
	tests := []struct {
		name    string
		d       *DSE
		args    args
		want    []*DSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", NewDSE() ,args{SortByHighPrice, ASC}, make([]*DSEShare,359), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DSE{}
			got, err := d.GetLatestPrices(tt.args.by, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("DSE.GetLatestPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("DSE.GetLatestPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDse(t *testing.T) {
	type args struct {
		arr   []*DSEShare
		by    sortBy
		order sortOrder
	}
	tests := []struct {
		name    string
		args    args
		want    []*DSEShare
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{make([]*DSEShare,0), SortByTradingCode, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByTradingCode, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByHighPrice, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByHighPrice, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByLowPrice, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByLowPrice, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByLTP, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByLTP, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByNumberOfTrades, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByNumberOfTrades, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByPriceChange, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByPriceChange, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByYCP, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByYCP, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByValue, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByValue, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByVolumeOfShare, ASC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByVolumeOfShare, DESC}, make([]*DSEShare,0), false},
		{"", args{make([]*DSEShare,0), SortByPercentageChange, DESC}, nil, true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortDse(tt.args.arr, tt.args.by, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortDse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDse() = %v, want %v", got, tt.want)
			}
		})
	}
}
