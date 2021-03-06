// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "mult", Type: field.TypeFloat64, Default: 1},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_users_tags",
				Columns:    []*schema.Column{TagsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "icon", Type: field.TypeInt, Default: 0},
		{Name: "title", Type: field.TypeString, Size: 64},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "estimated", Type: field.TypeInt, Nullable: true},
		{Name: "importance", Type: field.TypeInt8, Default: 5},
		{Name: "urgency", Type: field.TypeInt8, Default: 5},
		{Name: "custom_mult", Type: field.TypeFloat64, Default: 1},
		{Name: "lo", Type: field.TypeFloat64},
		{Name: "hi", Type: field.TypeFloat64},
		{Name: "tag_id", Type: field.TypeInt, Nullable: true},
		{Name: "creator_id", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_tags_tasks",
				Columns:    []*schema.Column{TasksColumns[12]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tasks_users_tasks",
				Columns:    []*schema.Column{TasksColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TagsTable,
		TasksTable,
		UsersTable,
	}
)

func init() {
	TagsTable.ForeignKeys[0].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = TagsTable
	TasksTable.ForeignKeys[1].RefTable = UsersTable
}
