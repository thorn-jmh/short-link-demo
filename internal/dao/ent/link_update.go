// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-svc-tpl/internal/dao/ent/link"
	"go-svc-tpl/internal/dao/ent/predicate"
	"go-svc-tpl/internal/dao/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetShort sets the "short" field.
func (lu *LinkUpdate) SetShort(s string) *LinkUpdate {
	lu.mutation.SetShort(s)
	return lu
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableShort(s *string) *LinkUpdate {
	if s != nil {
		lu.SetShort(*s)
	}
	return lu
}

// ClearShort clears the value of the "short" field.
func (lu *LinkUpdate) ClearShort() *LinkUpdate {
	lu.mutation.ClearShort()
	return lu
}

// SetOrigin sets the "origin" field.
func (lu *LinkUpdate) SetOrigin(s string) *LinkUpdate {
	lu.mutation.SetOrigin(s)
	return lu
}

// SetComment sets the "comment" field.
func (lu *LinkUpdate) SetComment(s string) *LinkUpdate {
	lu.mutation.SetComment(s)
	return lu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableComment(s *string) *LinkUpdate {
	if s != nil {
		lu.SetComment(*s)
	}
	return lu
}

// SetStartTime sets the "start_time" field.
func (lu *LinkUpdate) SetStartTime(t time.Time) *LinkUpdate {
	lu.mutation.SetStartTime(t)
	return lu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableStartTime(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetStartTime(*t)
	}
	return lu
}

// ClearStartTime clears the value of the "start_time" field.
func (lu *LinkUpdate) ClearStartTime() *LinkUpdate {
	lu.mutation.ClearStartTime()
	return lu
}

// SetEndTime sets the "end_time" field.
func (lu *LinkUpdate) SetEndTime(t time.Time) *LinkUpdate {
	lu.mutation.SetEndTime(t)
	return lu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableEndTime(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetEndTime(*t)
	}
	return lu
}

// ClearEndTime clears the value of the "end_time" field.
func (lu *LinkUpdate) ClearEndTime() *LinkUpdate {
	lu.mutation.ClearEndTime()
	return lu
}

// SetActive sets the "active" field.
func (lu *LinkUpdate) SetActive(b bool) *LinkUpdate {
	lu.mutation.SetActive(b)
	return lu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableActive(b *bool) *LinkUpdate {
	if b != nil {
		lu.SetActive(*b)
	}
	return lu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lu *LinkUpdate) SetOwnerID(id uint) *LinkUpdate {
	lu.mutation.SetOwnerID(id)
	return lu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (lu *LinkUpdate) SetNillableOwnerID(id *uint) *LinkUpdate {
	if id != nil {
		lu = lu.SetOwnerID(*id)
	}
	return lu
}

// SetOwner sets the "owner" edge to the User entity.
func (lu *LinkUpdate) SetOwner(u *User) *LinkUpdate {
	return lu.SetOwnerID(u.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lu *LinkUpdate) ClearOwner() *LinkUpdate {
	lu.mutation.ClearOwner()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LinkUpdate) check() error {
	if v, ok := lu.mutation.Short(); ok {
		if err := link.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf(`ent: validator failed for field "Link.short": %w`, err)}
		}
	}
	return nil
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Short(); ok {
		_spec.SetField(link.FieldShort, field.TypeString, value)
	}
	if lu.mutation.ShortCleared() {
		_spec.ClearField(link.FieldShort, field.TypeString)
	}
	if value, ok := lu.mutation.Origin(); ok {
		_spec.SetField(link.FieldOrigin, field.TypeString, value)
	}
	if value, ok := lu.mutation.Comment(); ok {
		_spec.SetField(link.FieldComment, field.TypeString, value)
	}
	if value, ok := lu.mutation.StartTime(); ok {
		_spec.SetField(link.FieldStartTime, field.TypeTime, value)
	}
	if lu.mutation.StartTimeCleared() {
		_spec.ClearField(link.FieldStartTime, field.TypeTime)
	}
	if value, ok := lu.mutation.EndTime(); ok {
		_spec.SetField(link.FieldEndTime, field.TypeTime, value)
	}
	if lu.mutation.EndTimeCleared() {
		_spec.ClearField(link.FieldEndTime, field.TypeTime)
	}
	if value, ok := lu.mutation.Active(); ok {
		_spec.SetField(link.FieldActive, field.TypeBool, value)
	}
	if lu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetShort sets the "short" field.
func (luo *LinkUpdateOne) SetShort(s string) *LinkUpdateOne {
	luo.mutation.SetShort(s)
	return luo
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableShort(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetShort(*s)
	}
	return luo
}

// ClearShort clears the value of the "short" field.
func (luo *LinkUpdateOne) ClearShort() *LinkUpdateOne {
	luo.mutation.ClearShort()
	return luo
}

// SetOrigin sets the "origin" field.
func (luo *LinkUpdateOne) SetOrigin(s string) *LinkUpdateOne {
	luo.mutation.SetOrigin(s)
	return luo
}

// SetComment sets the "comment" field.
func (luo *LinkUpdateOne) SetComment(s string) *LinkUpdateOne {
	luo.mutation.SetComment(s)
	return luo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableComment(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetComment(*s)
	}
	return luo
}

// SetStartTime sets the "start_time" field.
func (luo *LinkUpdateOne) SetStartTime(t time.Time) *LinkUpdateOne {
	luo.mutation.SetStartTime(t)
	return luo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableStartTime(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetStartTime(*t)
	}
	return luo
}

// ClearStartTime clears the value of the "start_time" field.
func (luo *LinkUpdateOne) ClearStartTime() *LinkUpdateOne {
	luo.mutation.ClearStartTime()
	return luo
}

// SetEndTime sets the "end_time" field.
func (luo *LinkUpdateOne) SetEndTime(t time.Time) *LinkUpdateOne {
	luo.mutation.SetEndTime(t)
	return luo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableEndTime(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetEndTime(*t)
	}
	return luo
}

// ClearEndTime clears the value of the "end_time" field.
func (luo *LinkUpdateOne) ClearEndTime() *LinkUpdateOne {
	luo.mutation.ClearEndTime()
	return luo
}

// SetActive sets the "active" field.
func (luo *LinkUpdateOne) SetActive(b bool) *LinkUpdateOne {
	luo.mutation.SetActive(b)
	return luo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableActive(b *bool) *LinkUpdateOne {
	if b != nil {
		luo.SetActive(*b)
	}
	return luo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (luo *LinkUpdateOne) SetOwnerID(id uint) *LinkUpdateOne {
	luo.mutation.SetOwnerID(id)
	return luo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableOwnerID(id *uint) *LinkUpdateOne {
	if id != nil {
		luo = luo.SetOwnerID(*id)
	}
	return luo
}

// SetOwner sets the "owner" edge to the User entity.
func (luo *LinkUpdateOne) SetOwner(u *User) *LinkUpdateOne {
	return luo.SetOwnerID(u.ID)
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (luo *LinkUpdateOne) ClearOwner() *LinkUpdateOne {
	luo.mutation.ClearOwner()
	return luo
}

// Where appends a list predicates to the LinkUpdate builder.
func (luo *LinkUpdateOne) Where(ps ...predicate.Link) *LinkUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LinkUpdateOne) check() error {
	if v, ok := luo.mutation.Short(); ok {
		if err := link.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf(`ent: validator failed for field "Link.short": %w`, err)}
		}
	}
	return nil
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Short(); ok {
		_spec.SetField(link.FieldShort, field.TypeString, value)
	}
	if luo.mutation.ShortCleared() {
		_spec.ClearField(link.FieldShort, field.TypeString)
	}
	if value, ok := luo.mutation.Origin(); ok {
		_spec.SetField(link.FieldOrigin, field.TypeString, value)
	}
	if value, ok := luo.mutation.Comment(); ok {
		_spec.SetField(link.FieldComment, field.TypeString, value)
	}
	if value, ok := luo.mutation.StartTime(); ok {
		_spec.SetField(link.FieldStartTime, field.TypeTime, value)
	}
	if luo.mutation.StartTimeCleared() {
		_spec.ClearField(link.FieldStartTime, field.TypeTime)
	}
	if value, ok := luo.mutation.EndTime(); ok {
		_spec.SetField(link.FieldEndTime, field.TypeTime, value)
	}
	if luo.mutation.EndTimeCleared() {
		_spec.ClearField(link.FieldEndTime, field.TypeTime)
	}
	if value, ok := luo.mutation.Active(); ok {
		_spec.SetField(link.FieldActive, field.TypeBool, value)
	}
	if luo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   link.OwnerTable,
			Columns: []string{link.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
