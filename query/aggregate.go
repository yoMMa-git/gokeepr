package query

type Aggregator interface {
	Calc(x float64)
	Total() (float64, error)
}
