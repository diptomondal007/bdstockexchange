package bdstockexchange

import (
	"testing"
)

func Test_isValidCategoryName(t *testing.T) {
	type args struct {
		categoryName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"",args{categoryName:"A"},true},
		{"",args{categoryName:"Abaa"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCategoryName(tt.args.categoryName); got != tt.want {
				t.Errorf("isValidCategoryName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeAmerican(t *testing.T) {
	type args struct {
		old string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"",args{old:"20,000.000"},"20000.000"},
		{"",args{old:"20,120"},"20120"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeAmerican(tt.args.old); got != tt.want {
				t.Errorf("normalizeAmerican() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFloat64(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"",args{text:"20,000.78"},20000.78},
		{"",args{text:"50.07"},50.07},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFloat64(tt.args.text); got != tt.want {
				t.Errorf("toFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toInt64(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"",args{text:"50,000"},50000},
		{"",args{text:"40,000"},40000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt64(tt.args.text); got != tt.want {
				t.Errorf("toInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toInt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"",args{text:"1"},1},
		{"",args{text:"2"},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt(tt.args.text); got != tt.want {
				t.Errorf("toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
