// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gmail-clone-backend/ent/email"
	"gmail-clone-backend/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailCreate is the builder for creating a Email entity.
type EmailCreate struct {
	config
	mutation *EmailMutation
	hooks    []Hook
}

// SetFrom sets the "from" field.
func (ec *EmailCreate) SetFrom(s string) *EmailCreate {
	ec.mutation.SetFrom(s)
	return ec
}

// SetSubject sets the "subject" field.
func (ec *EmailCreate) SetSubject(s string) *EmailCreate {
	ec.mutation.SetSubject(s)
	return ec
}

// SetBody sets the "body" field.
func (ec *EmailCreate) SetBody(s string) *EmailCreate {
	ec.mutation.SetBody(s)
	return ec
}

// SetSentAt sets the "sentAt" field.
func (ec *EmailCreate) SetSentAt(t time.Time) *EmailCreate {
	ec.mutation.SetSentAt(t)
	return ec
}

// SetArchived sets the "archived" field.
func (ec *EmailCreate) SetArchived(b bool) *EmailCreate {
	ec.mutation.SetArchived(b)
	return ec
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (ec *EmailCreate) SetNillableArchived(b *bool) *EmailCreate {
	if b != nil {
		ec.SetArchived(*b)
	}
	return ec
}

// SetRead sets the "read" field.
func (ec *EmailCreate) SetRead(b bool) *EmailCreate {
	ec.mutation.SetRead(b)
	return ec
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (ec *EmailCreate) SetNillableRead(b *bool) *EmailCreate {
	if b != nil {
		ec.SetRead(*b)
	}
	return ec
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ec *EmailCreate) SetUserID(id uuid.UUID) *EmailCreate {
	ec.mutation.SetUserID(id)
	return ec
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ec *EmailCreate) SetNillableUserID(id *uuid.UUID) *EmailCreate {
	if id != nil {
		ec = ec.SetUserID(*id)
	}
	return ec
}

// SetUser sets the "user" edge to the User entity.
func (ec *EmailCreate) SetUser(u *User) *EmailCreate {
	return ec.SetUserID(u.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (ec *EmailCreate) Mutation() *EmailMutation {
	return ec.mutation
}

// Save creates the Email in the database.
func (ec *EmailCreate) Save(ctx context.Context) (*Email, error) {
	var (
		err  error
		node *Email
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Email)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EmailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmailCreate) SaveX(ctx context.Context) *Email {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EmailCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EmailCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EmailCreate) defaults() {
	if _, ok := ec.mutation.Archived(); !ok {
		v := email.DefaultArchived
		ec.mutation.SetArchived(v)
	}
	if _, ok := ec.mutation.Read(); !ok {
		v := email.DefaultRead
		ec.mutation.SetRead(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EmailCreate) check() error {
	if _, ok := ec.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Email.from"`)}
	}
	if _, ok := ec.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "Email.subject"`)}
	}
	if _, ok := ec.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Email.body"`)}
	}
	if _, ok := ec.mutation.SentAt(); !ok {
		return &ValidationError{Name: "sentAt", err: errors.New(`ent: missing required field "Email.sentAt"`)}
	}
	if _, ok := ec.mutation.Archived(); !ok {
		return &ValidationError{Name: "archived", err: errors.New(`ent: missing required field "Email.archived"`)}
	}
	if _, ok := ec.mutation.Read(); !ok {
		return &ValidationError{Name: "read", err: errors.New(`ent: missing required field "Email.read"`)}
	}
	return nil
}

func (ec *EmailCreate) sqlSave(ctx context.Context) (*Email, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EmailCreate) createSpec() (*Email, *sqlgraph.CreateSpec) {
	var (
		_node = &Email{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: email.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: email.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: email.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := ec.mutation.Subject(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: email.FieldSubject,
		})
		_node.Subject = value
	}
	if value, ok := ec.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: email.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := ec.mutation.SentAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: email.FieldSentAt,
		})
		_node.SentAt = value
	}
	if value, ok := ec.mutation.Archived(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: email.FieldArchived,
		})
		_node.Archived = value
	}
	if value, ok := ec.mutation.Read(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: email.FieldRead,
		})
		_node.Read = value
	}
	if nodes := ec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.UserTable,
			Columns: []string{email.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_emails = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmailCreateBulk is the builder for creating many Email entities in bulk.
type EmailCreateBulk struct {
	config
	builders []*EmailCreate
}

// Save creates the Email entities in the database.
func (ecb *EmailCreateBulk) Save(ctx context.Context) ([]*Email, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Email, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EmailCreateBulk) SaveX(ctx context.Context) []*Email {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EmailCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EmailCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}