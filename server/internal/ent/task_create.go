// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/internal/ent/tag"
	"server/internal/ent/task"
	"server/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetIcon sets the "icon" field.
func (tc *TaskCreate) SetIcon(i int) *TaskCreate {
	tc.mutation.SetIcon(i)
	return tc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIcon(i *int) *TaskCreate {
	if i != nil {
		tc.SetIcon(*i)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *TaskCreate) SetTitle(s string) *TaskCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TaskCreate) SetDescription(s string) *TaskCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDescription(s *string) *TaskCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetDeadline sets the "deadline" field.
func (tc *TaskCreate) SetDeadline(t time.Time) *TaskCreate {
	tc.mutation.SetDeadline(t)
	return tc
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDeadline(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetDeadline(*t)
	}
	return tc
}

// SetEstimated sets the "estimated" field.
func (tc *TaskCreate) SetEstimated(i int) *TaskCreate {
	tc.mutation.SetEstimated(i)
	return tc
}

// SetNillableEstimated sets the "estimated" field if the given value is not nil.
func (tc *TaskCreate) SetNillableEstimated(i *int) *TaskCreate {
	if i != nil {
		tc.SetEstimated(*i)
	}
	return tc
}

// SetImportance sets the "importance" field.
func (tc *TaskCreate) SetImportance(i int8) *TaskCreate {
	tc.mutation.SetImportance(i)
	return tc
}

// SetNillableImportance sets the "importance" field if the given value is not nil.
func (tc *TaskCreate) SetNillableImportance(i *int8) *TaskCreate {
	if i != nil {
		tc.SetImportance(*i)
	}
	return tc
}

// SetUrgency sets the "urgency" field.
func (tc *TaskCreate) SetUrgency(i int8) *TaskCreate {
	tc.mutation.SetUrgency(i)
	return tc
}

// SetNillableUrgency sets the "urgency" field if the given value is not nil.
func (tc *TaskCreate) SetNillableUrgency(i *int8) *TaskCreate {
	if i != nil {
		tc.SetUrgency(*i)
	}
	return tc
}

// SetCustomMult sets the "custom_mult" field.
func (tc *TaskCreate) SetCustomMult(f float64) *TaskCreate {
	tc.mutation.SetCustomMult(f)
	return tc
}

// SetNillableCustomMult sets the "custom_mult" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCustomMult(f *float64) *TaskCreate {
	if f != nil {
		tc.SetCustomMult(*f)
	}
	return tc
}

// SetLo sets the "lo" field.
func (tc *TaskCreate) SetLo(f float64) *TaskCreate {
	tc.mutation.SetLo(f)
	return tc
}

// SetHi sets the "hi" field.
func (tc *TaskCreate) SetHi(f float64) *TaskCreate {
	tc.mutation.SetHi(f)
	return tc
}

// SetTagID sets the "tag_id" field.
func (tc *TaskCreate) SetTagID(i int) *TaskCreate {
	tc.mutation.SetTagID(i)
	return tc
}

// SetNillableTagID sets the "tag_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableTagID(i *int) *TaskCreate {
	if i != nil {
		tc.SetTagID(*i)
	}
	return tc
}

// SetCreatorID sets the "creator_id" field.
func (tc *TaskCreate) SetCreatorID(i int) *TaskCreate {
	tc.mutation.SetCreatorID(i)
	return tc
}

// SetCreator sets the "creator" edge to the User entity.
func (tc *TaskCreate) SetCreator(u *User) *TaskCreate {
	return tc.SetCreatorID(u.ID)
}

// SetTaggID sets the "tagg" edge to the Tag entity by ID.
func (tc *TaskCreate) SetTaggID(id int) *TaskCreate {
	tc.mutation.SetTaggID(id)
	return tc
}

// SetNillableTaggID sets the "tagg" edge to the Tag entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableTaggID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetTaggID(*id)
	}
	return tc
}

// SetTagg sets the "tagg" edge to the Tag entity.
func (tc *TaskCreate) SetTagg(t *Tag) *TaskCreate {
	return tc.SetTaggID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.Icon(); !ok {
		v := task.DefaultIcon
		tc.mutation.SetIcon(v)
	}
	if _, ok := tc.mutation.Importance(); !ok {
		v := task.DefaultImportance
		tc.mutation.SetImportance(v)
	}
	if _, ok := tc.mutation.Urgency(); !ok {
		v := task.DefaultUrgency
		tc.mutation.SetUrgency(v)
	}
	if _, ok := tc.mutation.CustomMult(); !ok {
		v := task.DefaultCustomMult
		tc.mutation.SetCustomMult(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "icon"`)}
	}
	if v, ok := tc.mutation.Icon(); ok {
		if err := task.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "icon": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := tc.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Description(); ok {
		if err := task.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	if v, ok := tc.mutation.Estimated(); ok {
		if err := task.EstimatedValidator(v); err != nil {
			return &ValidationError{Name: "estimated", err: fmt.Errorf(`ent: validator failed for field "estimated": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Importance(); !ok {
		return &ValidationError{Name: "importance", err: errors.New(`ent: missing required field "importance"`)}
	}
	if v, ok := tc.mutation.Importance(); ok {
		if err := task.ImportanceValidator(v); err != nil {
			return &ValidationError{Name: "importance", err: fmt.Errorf(`ent: validator failed for field "importance": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Urgency(); !ok {
		return &ValidationError{Name: "urgency", err: errors.New(`ent: missing required field "urgency"`)}
	}
	if v, ok := tc.mutation.Urgency(); ok {
		if err := task.UrgencyValidator(v); err != nil {
			return &ValidationError{Name: "urgency", err: fmt.Errorf(`ent: validator failed for field "urgency": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CustomMult(); !ok {
		return &ValidationError{Name: "custom_mult", err: errors.New(`ent: missing required field "custom_mult"`)}
	}
	if _, ok := tc.mutation.Lo(); !ok {
		return &ValidationError{Name: "lo", err: errors.New(`ent: missing required field "lo"`)}
	}
	if _, ok := tc.mutation.Hi(); !ok {
		return &ValidationError{Name: "hi", err: errors.New(`ent: missing required field "hi"`)}
	}
	if _, ok := tc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator_id", err: errors.New(`ent: missing required field "creator_id"`)}
	}
	if _, ok := tc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator", err: errors.New("ent: missing required edge \"creator\"")}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tc.mutation.Deadline(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDeadline,
		})
		_node.Deadline = value
	}
	if value, ok := tc.mutation.Estimated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: task.FieldEstimated,
		})
		_node.Estimated = value
	}
	if value, ok := tc.mutation.Importance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldImportance,
		})
		_node.Importance = value
	}
	if value, ok := tc.mutation.Urgency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: task.FieldUrgency,
		})
		_node.Urgency = value
	}
	if value, ok := tc.mutation.CustomMult(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: task.FieldCustomMult,
		})
		_node.CustomMult = value
	}
	if value, ok := tc.mutation.Lo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: task.FieldLo,
		})
		_node.Lo = value
	}
	if value, ok := tc.mutation.Hi(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: task.FieldHi,
		})
		_node.Hi = value
	}
	if nodes := tc.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_node.CreatorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TaggIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TaggTable,
			Columns: []string{task.TaggColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
