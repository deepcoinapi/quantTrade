package data

type Tick struct {
	Symbol      string
	SpotPrice   float64
	FuturePrice float64
}

type Bar struct {
	Ts     int64
	Symbol string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}
