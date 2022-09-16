// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/haikara-dev/haikara/ent/predicate"
	"github.com/haikara-dev/haikara/ent/sitecrawlrule"
)

// SiteCrawlRuleDelete is the builder for deleting a SiteCrawlRule entity.
type SiteCrawlRuleDelete struct {
	config
	hooks    []Hook
	mutation *SiteCrawlRuleMutation
}

// Where appends a list predicates to the SiteCrawlRuleDelete builder.
func (scrd *SiteCrawlRuleDelete) Where(ps ...predicate.SiteCrawlRule) *SiteCrawlRuleDelete {
	scrd.mutation.Where(ps...)
	return scrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scrd *SiteCrawlRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scrd.hooks) == 0 {
		affected, err = scrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteCrawlRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scrd.mutation = mutation
			affected, err = scrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scrd.hooks) - 1; i >= 0; i-- {
			if scrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scrd *SiteCrawlRuleDelete) ExecX(ctx context.Context) int {
	n, err := scrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scrd *SiteCrawlRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sitecrawlrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sitecrawlrule.FieldID,
			},
		},
	}
	if ps := scrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SiteCrawlRuleDeleteOne is the builder for deleting a single SiteCrawlRule entity.
type SiteCrawlRuleDeleteOne struct {
	scrd *SiteCrawlRuleDelete
}

// Exec executes the deletion query.
func (scrdo *SiteCrawlRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := scrdo.scrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sitecrawlrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scrdo *SiteCrawlRuleDeleteOne) ExecX(ctx context.Context) {
	scrdo.scrd.ExecX(ctx)
}
