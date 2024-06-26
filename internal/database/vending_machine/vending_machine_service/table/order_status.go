//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var OrderStatus = newOrderStatusTable("vending_machine_service", "order_status", "")

type orderStatusTable struct {
	postgres.Table

	// Columns
	ID      postgres.ColumnInteger
	OrderID postgres.ColumnInteger
	Status  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrderStatusTable struct {
	orderStatusTable

	EXCLUDED orderStatusTable
}

// AS creates new OrderStatusTable with assigned alias
func (a OrderStatusTable) AS(alias string) *OrderStatusTable {
	return newOrderStatusTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrderStatusTable with assigned schema name
func (a OrderStatusTable) FromSchema(schemaName string) *OrderStatusTable {
	return newOrderStatusTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrderStatusTable with assigned table prefix
func (a OrderStatusTable) WithPrefix(prefix string) *OrderStatusTable {
	return newOrderStatusTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrderStatusTable with assigned table suffix
func (a OrderStatusTable) WithSuffix(suffix string) *OrderStatusTable {
	return newOrderStatusTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrderStatusTable(schemaName, tableName, alias string) *OrderStatusTable {
	return &OrderStatusTable{
		orderStatusTable: newOrderStatusTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newOrderStatusTableImpl("", "excluded", ""),
	}
}

func newOrderStatusTableImpl(schemaName, tableName, alias string) orderStatusTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		OrderIDColumn  = postgres.IntegerColumn("order_id")
		StatusColumn   = postgres.StringColumn("status")
		allColumns     = postgres.ColumnList{IDColumn, OrderIDColumn, StatusColumn}
		mutableColumns = postgres.ColumnList{OrderIDColumn, StatusColumn}
	)

	return orderStatusTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		OrderID: OrderIDColumn,
		Status:  StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
