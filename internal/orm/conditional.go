package orm

const (
	ActionEqual = "EQUAL"
)

type Wheres struct {
	where  string
	values []any
}

func NewWheres(where string, values ...any) Wheres {
	return Wheres{where: where, values: values}
}
