// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cubdesign/haikara/ent/predicate"
	"github.com/cubdesign/haikara/ent/site"
	"github.com/cubdesign/haikara/ent/sitecategory"
)

// SiteCategoryUpdate is the builder for updating SiteCategory entities.
type SiteCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *SiteCategoryMutation
}

// Where appends a list predicates to the SiteCategoryUpdate builder.
func (scu *SiteCategoryUpdate) Where(ps ...predicate.SiteCategory) *SiteCategoryUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *SiteCategoryUpdate) SetUpdatedAt(t time.Time) *SiteCategoryUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// SetLabel sets the "label" field.
func (scu *SiteCategoryUpdate) SetLabel(s string) *SiteCategoryUpdate {
	scu.mutation.SetLabel(s)
	return scu
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (scu *SiteCategoryUpdate) AddSiteIDs(ids ...int) *SiteCategoryUpdate {
	scu.mutation.AddSiteIDs(ids...)
	return scu
}

// AddSites adds the "sites" edges to the Site entity.
func (scu *SiteCategoryUpdate) AddSites(s ...*Site) *SiteCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.AddSiteIDs(ids...)
}

// Mutation returns the SiteCategoryMutation object of the builder.
func (scu *SiteCategoryUpdate) Mutation() *SiteCategoryMutation {
	return scu.mutation
}

// ClearSites clears all "sites" edges to the Site entity.
func (scu *SiteCategoryUpdate) ClearSites() *SiteCategoryUpdate {
	scu.mutation.ClearSites()
	return scu
}

// RemoveSiteIDs removes the "sites" edge to Site entities by IDs.
func (scu *SiteCategoryUpdate) RemoveSiteIDs(ids ...int) *SiteCategoryUpdate {
	scu.mutation.RemoveSiteIDs(ids...)
	return scu
}

// RemoveSites removes "sites" edges to Site entities.
func (scu *SiteCategoryUpdate) RemoveSites(s ...*Site) *SiteCategoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scu.RemoveSiteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SiteCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	scu.defaults()
	if len(scu.hooks) == 0 {
		if err = scu.check(); err != nil {
			return 0, err
		}
		affected, err = scu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scu.check(); err != nil {
				return 0, err
			}
			scu.mutation = mutation
			affected, err = scu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scu.hooks) - 1; i >= 0; i-- {
			if scu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SiteCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SiteCategoryUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SiteCategoryUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *SiteCategoryUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := sitecategory.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *SiteCategoryUpdate) check() error {
	if v, ok := scu.mutation.Label(); ok {
		if err := sitecategory.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "SiteCategory.label": %w`, err)}
		}
	}
	return nil
}

func (scu *SiteCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sitecategory.Table,
			Columns: sitecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sitecategory.FieldID,
			},
		},
	}
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sitecategory.FieldUpdatedAt,
		})
	}
	if value, ok := scu.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sitecategory.FieldLabel,
		})
	}
	if scu.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedSitesIDs(); len(nodes) > 0 && !scu.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sitecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SiteCategoryUpdateOne is the builder for updating a single SiteCategory entity.
type SiteCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SiteCategoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *SiteCategoryUpdateOne) SetUpdatedAt(t time.Time) *SiteCategoryUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// SetLabel sets the "label" field.
func (scuo *SiteCategoryUpdateOne) SetLabel(s string) *SiteCategoryUpdateOne {
	scuo.mutation.SetLabel(s)
	return scuo
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (scuo *SiteCategoryUpdateOne) AddSiteIDs(ids ...int) *SiteCategoryUpdateOne {
	scuo.mutation.AddSiteIDs(ids...)
	return scuo
}

// AddSites adds the "sites" edges to the Site entity.
func (scuo *SiteCategoryUpdateOne) AddSites(s ...*Site) *SiteCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.AddSiteIDs(ids...)
}

// Mutation returns the SiteCategoryMutation object of the builder.
func (scuo *SiteCategoryUpdateOne) Mutation() *SiteCategoryMutation {
	return scuo.mutation
}

// ClearSites clears all "sites" edges to the Site entity.
func (scuo *SiteCategoryUpdateOne) ClearSites() *SiteCategoryUpdateOne {
	scuo.mutation.ClearSites()
	return scuo
}

// RemoveSiteIDs removes the "sites" edge to Site entities by IDs.
func (scuo *SiteCategoryUpdateOne) RemoveSiteIDs(ids ...int) *SiteCategoryUpdateOne {
	scuo.mutation.RemoveSiteIDs(ids...)
	return scuo
}

// RemoveSites removes "sites" edges to Site entities.
func (scuo *SiteCategoryUpdateOne) RemoveSites(s ...*Site) *SiteCategoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return scuo.RemoveSiteIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SiteCategoryUpdateOne) Select(field string, fields ...string) *SiteCategoryUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SiteCategory entity.
func (scuo *SiteCategoryUpdateOne) Save(ctx context.Context) (*SiteCategory, error) {
	var (
		err  error
		node *SiteCategory
	)
	scuo.defaults()
	if len(scuo.hooks) == 0 {
		if err = scuo.check(); err != nil {
			return nil, err
		}
		node, err = scuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scuo.check(); err != nil {
				return nil, err
			}
			scuo.mutation = mutation
			node, err = scuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(scuo.hooks) - 1; i >= 0; i-- {
			if scuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, scuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SiteCategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SiteCategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SiteCategoryUpdateOne) SaveX(ctx context.Context) *SiteCategory {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SiteCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SiteCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *SiteCategoryUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := sitecategory.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *SiteCategoryUpdateOne) check() error {
	if v, ok := scuo.mutation.Label(); ok {
		if err := sitecategory.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "SiteCategory.label": %w`, err)}
		}
	}
	return nil
}

func (scuo *SiteCategoryUpdateOne) sqlSave(ctx context.Context) (_node *SiteCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sitecategory.Table,
			Columns: sitecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sitecategory.FieldID,
			},
		},
	}
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SiteCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sitecategory.FieldID)
		for _, f := range fields {
			if !sitecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sitecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sitecategory.FieldUpdatedAt,
		})
	}
	if value, ok := scuo.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sitecategory.FieldLabel,
		})
	}
	if scuo.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedSitesIDs(); len(nodes) > 0 && !scuo.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sitecategory.SitesTable,
			Columns: sitecategory.SitesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SiteCategory{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sitecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
