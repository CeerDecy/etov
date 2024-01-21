package orm

import "gorm.io/gorm"

const (
	ActionAnd = "AND"
	ActionOr  = "OR"
)

type Wheres struct {
	where  string
	values []any
	action string
}

func And(where string, values ...any) Wheres {
	return Wheres{where: where, values: values, action: ActionAnd}
}

func Or(where string, values ...any) Wheres {
	return Wheres{where: where, values: values, action: ActionOr}
}

func MountTransaction(tx *gorm.DB, wheres ...Wheres) *gorm.DB {
	for _, where := range wheres {
		if where.action == ActionOr {
			tx.Or(where.where, where.values...)
			continue
		}
		tx.Where(where.where, where.values...)
	}
	return tx
}
