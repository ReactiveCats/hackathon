// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/internal/ent/task"
	"server/internal/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon int `json:"icon,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline time.Time `json:"deadline,omitempty"`
	// Estimated holds the value of the "estimated" field.
	Estimated int `json:"estimated,omitempty"`
	// Complexity holds the value of the "complexity" field.
	Complexity int8 `json:"complexity,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int8 `json:"priority,omitempty"`
	// F holds the value of the "f" field.
	F float64 `json:"f,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID int `json:"creator_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges TaskEdges `json:"edges"`
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// The edge creator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldF:
			values[i] = new(sql.NullFloat64)
		case task.FieldID, task.FieldIcon, task.FieldEstimated, task.FieldComplexity, task.FieldPriority, task.FieldCreatorID:
			values[i] = new(sql.NullInt64)
		case task.FieldTitle, task.FieldDescription:
			values[i] = new(sql.NullString)
		case task.FieldCreatedAt, task.FieldDeadline:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldIcon:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				t.Icon = int(value.Int64)
			}
		case task.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case task.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case task.FieldDeadline:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				t.Deadline = value.Time
			}
		case task.FieldEstimated:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field estimated", values[i])
			} else if value.Valid {
				t.Estimated = int(value.Int64)
			}
		case task.FieldComplexity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field complexity", values[i])
			} else if value.Valid {
				t.Complexity = int8(value.Int64)
			}
		case task.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int8(value.Int64)
			}
		case task.FieldF:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field f", values[i])
			} else if value.Valid {
				t.F = value.Float64
			}
		case task.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				t.CreatorID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCreator queries the "creator" edge of the Task entity.
func (t *Task) QueryCreator() *UserQuery {
	return (&TaskClient{config: t.config}).QueryCreator(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", icon=")
	builder.WriteString(fmt.Sprintf("%v", t.Icon))
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", deadline=")
	builder.WriteString(t.Deadline.Format(time.ANSIC))
	builder.WriteString(", estimated=")
	builder.WriteString(fmt.Sprintf("%v", t.Estimated))
	builder.WriteString(", complexity=")
	builder.WriteString(fmt.Sprintf("%v", t.Complexity))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", f=")
	builder.WriteString(fmt.Sprintf("%v", t.F))
	builder.WriteString(", creator_id=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatorID))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
