package common

import (
	"fmt"
	"math"
	"strconv"
)

func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//四舍五入
func Float64ToInt(v float64) int {
	return int(math.Round(v))
}

//底价根据利润率换算, 保留小数点后2位
func ConvertDevPrice2DspPrice(price float64, rate float64) float64 {
	if rate == 0 {
		return price
	}
	price, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", price / (1- rate / 100)), 64)
	return price
}

//价格换算dsp -> dev, 保留小数点后7位
func ConvertDspPrice2DevPrice(price float64, rate float64) float64 {
	if rate == 0 {
		return price
	}
	price, _ = strconv.ParseFloat(fmt.Sprintf("%.7f", price * (1- rate / 100)), 64)
	return price
}
