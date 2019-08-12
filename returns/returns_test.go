package main

import (
	"math"
	"reflect"
	"testing"
)

func TestPriceMultiPeriodSimpleReturn(t *testing.T) {
	type args struct {
		ts TimeSeries
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"case1", args{TimeSeries{6, []float64{100, 120, 98, 130, 140, 190, 200}}}, 0.122462048},
		{"case2", args{TimeSeries{1, []float64{100, 120}}}, 0.2},
		{"case3", args{TimeSeries{2, []float64{100, 50, 300}}}, 0.7320508},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PriceMultiPeriodSimpleReturn(tt.args.ts); math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("PriceMultiPeriodSimpleReturn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnePeriodSimpleReturn(t *testing.T) {
	type args struct {
		ts TimeSeries
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnePeriodSimpleReturn(tt.args.ts); got != tt.want {
				t.Errorf("OnePeriodSimpleReturn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPMPSR2(t *testing.T) {
	type args struct {
		ts TimeSeries
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PMPSR2(tt.args.ts); got != tt.want {
				t.Errorf("PMPSR2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRk(t *testing.T) {
	type args struct {
		ts TimeSeries
	}
	tests := []struct {
		name string
		args args
		want TimeSeries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rk(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
