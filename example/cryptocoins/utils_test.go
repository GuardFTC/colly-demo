// @Author:冯铁城 [17615007230@163.com] 2025-10-21 16:49:47
package cryptocoins

import "testing"

func TestExtractFullAmount(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractFullAmount(tt.args.input); got != tt.want {
				t.Errorf("ExtractFullAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatPercent(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatPercent(tt.args.value); got != tt.want {
				t.Errorf("FormatPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUSD(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatUSD(tt.args.value); got != tt.want {
				t.Errorf("FormatUSD() = %v, want %v", got, tt.want)
			}
		})
	}
}
