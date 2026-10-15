package catalog

type ColumnType int

const (
	TypeText ColumnType = iota
	TypeInt
	TypeFloat
)

type Column struct {
	Name string
	Type ColumnType
}

// Предполагаемый ключ: <table>-<partition>-<clustering с незначимыми нулями>-<column>
type Schema struct {
	Table         string
	Columns       []Column
	PartitionKey  string
	ClusteringKey []string
}

type Catalog interface {
	CreateTable(s Schema) error
	GetTable(name string) (Schema, error)
	DropTable(name string) error
}
