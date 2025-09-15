package indicator

import (
	"math"
	"quantTrade/core/data"
)

// calcSMAStd: return SMA and population std dev
func calcSMAStd(vals []data.Bar) (float64, float64) {
	n := float64(len(vals))
	sum := 0.0
	for _, v := range vals {
		//fmt.Print(" close:", v.Close)
		sum += v.Close
	}
	//fmt.Println("===sum:", sum, n)
	ma := sum / n
	variance := 0.0
	for _, v := range vals {
		diff := v.Close - ma
		variance += diff * diff
	}
	std := math.Sqrt(variance / n)
	return ma, std
}
