// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/internal/ent/predicate"
	"server/internal/ent/task"
	"server/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetCreatorID sets the "creator_id" field.
func (tu *TaskUpdate) SetCreatorID(i int) *TaskUpdate {
	tu.mutation.SetCreatorID(i)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TaskUpdate) SetPriority(s string) *TaskUpdate {
	tu.mutation.SetPriority(s)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePriority(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPriority(*s)
	}
	return tu
}

// SetComplexity sets the "complexity" field.
func (tu *TaskUpdate) SetComplexity(s string) *TaskUpdate {
	tu.mutation.SetComplexity(s)
	return tu
}

// SetNillableComplexity sets the "complexity" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableComplexity(s *string) *TaskUpdate {
	if s != nil {
		tu.SetComplexity(*s)
	}
	return tu
}

// SetHardDeadline sets the "hard_deadline" field.
func (tu *TaskUpdate) SetHardDeadline(t time.Time) *TaskUpdate {
	tu.mutation.SetHardDeadline(t)
	return tu
}

// SetNillableHardDeadline sets the "hard_deadline" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableHardDeadline(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetHardDeadline(*t)
	}
	return tu
}

// ClearHardDeadline clears the value of the "hard_deadline" field.
func (tu *TaskUpdate) ClearHardDeadline() *TaskUpdate {
	tu.mutation.ClearHardDeadline()
	return tu
}

// SetSoftDeadline sets the "soft_deadline" field.
func (tu *TaskUpdate) SetSoftDeadline(t time.Time) *TaskUpdate {
	tu.mutation.SetSoftDeadline(t)
	return tu
}

// SetNillableSoftDeadline sets the "soft_deadline" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableSoftDeadline(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetSoftDeadline(*t)
	}
	return tu
}

// ClearSoftDeadline clears the value of the "soft_deadline" field.
func (tu *TaskUpdate) ClearSoftDeadline() *TaskUpdate {
	tu.mutation.ClearSoftDeadline()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(s string) *TaskUpdate {
	tu.mutation.SetStatus(s)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(s *string) *TaskUpdate {
	if s != nil {
		tu.SetStatus(*s)
	}
	return tu
}

// SetCreator sets the "creator" edge to the User entity.
func (tu *TaskUpdate) SetCreator(u *User) *TaskUpdate {
	return tu.SetCreatorID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearCreator clears the "creator" edge to the User entity.
func (tu *TaskUpdate) ClearCreator() *TaskUpdate {
	tu.mutation.ClearCreator()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if _, ok := tu.mutation.CreatorID(); tu.mutation.CreatorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"creator\"")
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPriority,
		})
	}
	if value, ok := tu.mutation.Complexity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldComplexity,
		})
	}
	if value, ok := tu.mutation.HardDeadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldHardDeadline,
		})
	}
	if tu.mutation.HardDeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldHardDeadline,
		})
	}
	if value, ok := tu.mutation.SoftDeadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldSoftDeadline,
		})
	}
	if tu.mutation.SoftDeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldSoftDeadline,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStatus,
		})
	}
	if tu.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CreatorTable,
			Columns: []string{task.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CreatorTable,
			Columns: []string{task.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetCreatorID sets the "creator_id" field.
func (tuo *TaskUpdateOne) SetCreatorID(i int) *TaskUpdateOne {
	tuo.mutation.SetCreatorID(i)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TaskUpdateOne) SetPriority(s string) *TaskUpdateOne {
	tuo.mutation.SetPriority(s)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePriority(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPriority(*s)
	}
	return tuo
}

// SetComplexity sets the "complexity" field.
func (tuo *TaskUpdateOne) SetComplexity(s string) *TaskUpdateOne {
	tuo.mutation.SetComplexity(s)
	return tuo
}

// SetNillableComplexity sets the "complexity" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableComplexity(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetComplexity(*s)
	}
	return tuo
}

// SetHardDeadline sets the "hard_deadline" field.
func (tuo *TaskUpdateOne) SetHardDeadline(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetHardDeadline(t)
	return tuo
}

// SetNillableHardDeadline sets the "hard_deadline" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableHardDeadline(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetHardDeadline(*t)
	}
	return tuo
}

// ClearHardDeadline clears the value of the "hard_deadline" field.
func (tuo *TaskUpdateOne) ClearHardDeadline() *TaskUpdateOne {
	tuo.mutation.ClearHardDeadline()
	return tuo
}

// SetSoftDeadline sets the "soft_deadline" field.
func (tuo *TaskUpdateOne) SetSoftDeadline(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetSoftDeadline(t)
	return tuo
}

// SetNillableSoftDeadline sets the "soft_deadline" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableSoftDeadline(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetSoftDeadline(*t)
	}
	return tuo
}

// ClearSoftDeadline clears the value of the "soft_deadline" field.
func (tuo *TaskUpdateOne) ClearSoftDeadline() *TaskUpdateOne {
	tuo.mutation.ClearSoftDeadline()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(s string) *TaskUpdateOne {
	tuo.mutation.SetStatus(s)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetStatus(*s)
	}
	return tuo
}

// SetCreator sets the "creator" edge to the User entity.
func (tuo *TaskUpdateOne) SetCreator(u *User) *TaskUpdateOne {
	return tuo.SetCreatorID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearCreator clears the "creator" edge to the User entity.
func (tuo *TaskUpdateOne) ClearCreator() *TaskUpdateOne {
	tuo.mutation.ClearCreator()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if _, ok := tuo.mutation.CreatorID(); tuo.mutation.CreatorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"creator\"")
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldPriority,
		})
	}
	if value, ok := tuo.mutation.Complexity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldComplexity,
		})
	}
	if value, ok := tuo.mutation.HardDeadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldHardDeadline,
		})
	}
	if tuo.mutation.HardDeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldHardDeadline,
		})
	}
	if value, ok := tuo.mutation.SoftDeadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldSoftDeadline,
		})
	}
	if tuo.mutation.SoftDeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldSoftDeadline,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldStatus,
		})
	}
	if tuo.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CreatorTable,
			Columns: []string{task.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CreatorTable,
			Columns: []string{task.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
