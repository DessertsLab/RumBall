package main

import (
	"fmt"
	"math"
)

type TimeSeries struct {
	period int
	ts     []float64
}

/*
OnePeriodSimpleReturn 单期简单收益率
输入资产价格序列
*/
func OnePeriodSimpleReturn(ts TimeSeries) float64 {
	if ts.period != 2 {
		panic("ts length should be 2")
	}
	return ts.ts[1]/ts.ts[0] - 1
}

/*
PriceMultiPeriodSimpleReturn 多期简单收益率 输入每期收益率序列
输入资产价格序列
输入期初期末价格

返回期(年）化收益率 double
每期收益率序列

*/
func PriceMultiPeriodSimpleReturn(ts TimeSeries) float64 {
	var r float64
	for i := 0; i < ts.period; i++ {
		rk := ts.ts[i+1]/ts.ts[i] - 1
		r += math.Log(1 + rk)
	}

	return math.Exp(r/float64(ts.period)) - 1
}

func PMPSR2(ts TimeSeries) float64 {
	return math.Pow(ts.ts[ts.period]/ts.ts[0], 1/float64(ts.period)) - 1
}

/*
Rk 每期的收益率序列
*/
func Rk(ts TimeSeries) TimeSeries {
	for i := 0; i < ts.period-1; i++ {
		ts.ts[i] = Round(ts.ts[i+1]/ts.ts[i]-1, 6)
	}

	return TimeSeries{ts.period - 1, ts.ts[:ts.period-1]}
}

func Round(a float64, b float64) float64 {
	return math.Round(a*(math.Pow(10, b))) / math.Pow(10, b)
}

func main() {
	r1 := OnePeriodSimpleReturn(TimeSeries{2, []float64{100, 150, 1000}})
	fmt.Println(r1)
	r2 := PriceMultiPeriodSimpleReturn(TimeSeries{6, []float64{100, 120, 98, 130, 140, 190, 200}})
	fmt.Println(r2)
	r3 := Rk(TimeSeries{6, []float64{100, 120, 98, 130, 140, 190, 200}})
	fmt.Println(r3)
	r4 := PMPSR2(TimeSeries{6, []float64{100, 120, 98, 130, 140, 190, 200}})
	fmt.Println(r4)
	r5 := PriceMultiPeriodSimpleReturn(TimeSeries{1, []float64{100, 120}})
	fmt.Println(r5)
}
