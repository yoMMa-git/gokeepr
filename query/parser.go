package query

type AggFunc string

const (
	AggNone  AggFunc = ""
	AggCount AggFunc = "COUNT"
	AggSum   AggFunc = "SUM"
	AggMin   AggFunc = "MIN"
	AggMax   AggFunc = "MAX"
	AggAvg   AggFunc = "AVG"
)

type Statement struct {
	Kind    string // "CREATE", "INSERT", "SELECT", "DELETE"
	Table   string
	Columns []string
	Values  []string
	Where   map[string]string // фильтрация по ключевым колонкам
	Agg     AggFunc
	AggCol  string
}

// CQL -> Statement
type Parser interface {
	Parse(query string) (Statement, error)
}
