package query

type Result struct {
	Rows  []map[string]any
	Value float64 // для агрегаций
}

// FROM -> схема из каталога; WHERE по ключу -> границы Scan; агрегация -> Aggregator.Calc на каждое значение, затем Total.
type Executor interface {
	Execute(st Statement) (Result, error)
}
