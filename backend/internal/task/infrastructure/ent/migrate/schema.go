// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 20},
		{Name: "description", Type: field.TypeString, Size: 255},
		{Name: "priority", Type: field.TypeEnum, Enums: []string{"LOW", "MIDDLE", "HIGH"}, Default: "MIDDLE"},
		{Name: "stage", Type: field.TypeEnum, Enums: []string{"TODO", "IN_PROGRESS", "DONE"}},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TasksTable,
	}
)

func init() {
}
