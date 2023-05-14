// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/ent/conversations"
)

// ConversationsCreate is the builder for creating a Conversations entity.
type ConversationsCreate struct {
	config
	mutation *ConversationsMutation
	hooks    []Hook
}

// SetAiModel sets the "ai_model" field.
func (cc *ConversationsCreate) SetAiModel(cm conversations.AiModel) *ConversationsCreate {
	cc.mutation.SetAiModel(cm)
	return cc
}

// SetSnsType sets the "sns_type" field.
func (cc *ConversationsCreate) SetSnsType(ct conversations.SnsType) *ConversationsCreate {
	cc.mutation.SetSnsType(ct)
	return cc
}

// SetCmdVersion sets the "cmd_version" field.
func (cc *ConversationsCreate) SetCmdVersion(cv conversations.CmdVersion) *ConversationsCreate {
	cc.mutation.SetCmdVersion(cv)
	return cc
}

// SetIsAborted sets the "is_aborted" field.
func (cc *ConversationsCreate) SetIsAborted(b bool) *ConversationsCreate {
	cc.mutation.SetIsAborted(b)
	return cc
}

// SetNillableIsAborted sets the "is_aborted" field if the given value is not nil.
func (cc *ConversationsCreate) SetNillableIsAborted(b *bool) *ConversationsCreate {
	if b != nil {
		cc.SetIsAborted(*b)
	}
	return cc
}

// SetAbortReason sets the "abort_reason" field.
func (cc *ConversationsCreate) SetAbortReason(s string) *ConversationsCreate {
	cc.mutation.SetAbortReason(s)
	return cc
}

// SetNillableAbortReason sets the "abort_reason" field if the given value is not nil.
func (cc *ConversationsCreate) SetNillableAbortReason(s *string) *ConversationsCreate {
	if s != nil {
		cc.SetAbortReason(*s)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ConversationsCreate) SetCreatedAt(t time.Time) *ConversationsCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ConversationsCreate) SetNillableCreatedAt(t *time.Time) *ConversationsCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// Mutation returns the ConversationsMutation object of the builder.
func (cc *ConversationsCreate) Mutation() *ConversationsMutation {
	return cc.mutation
}

// Save creates the Conversations in the database.
func (cc *ConversationsCreate) Save(ctx context.Context) (*Conversations, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConversationsCreate) SaveX(ctx context.Context) *Conversations {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConversationsCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConversationsCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ConversationsCreate) defaults() {
	if _, ok := cc.mutation.IsAborted(); !ok {
		v := conversations.DefaultIsAborted
		cc.mutation.SetIsAborted(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := conversations.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConversationsCreate) check() error {
	if _, ok := cc.mutation.AiModel(); !ok {
		return &ValidationError{Name: "ai_model", err: errors.New(`ent: missing required field "Conversations.ai_model"`)}
	}
	if v, ok := cc.mutation.AiModel(); ok {
		if err := conversations.AiModelValidator(v); err != nil {
			return &ValidationError{Name: "ai_model", err: fmt.Errorf(`ent: validator failed for field "Conversations.ai_model": %w`, err)}
		}
	}
	if _, ok := cc.mutation.SnsType(); !ok {
		return &ValidationError{Name: "sns_type", err: errors.New(`ent: missing required field "Conversations.sns_type"`)}
	}
	if v, ok := cc.mutation.SnsType(); ok {
		if err := conversations.SnsTypeValidator(v); err != nil {
			return &ValidationError{Name: "sns_type", err: fmt.Errorf(`ent: validator failed for field "Conversations.sns_type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CmdVersion(); !ok {
		return &ValidationError{Name: "cmd_version", err: errors.New(`ent: missing required field "Conversations.cmd_version"`)}
	}
	if v, ok := cc.mutation.CmdVersion(); ok {
		if err := conversations.CmdVersionValidator(v); err != nil {
			return &ValidationError{Name: "cmd_version", err: fmt.Errorf(`ent: validator failed for field "Conversations.cmd_version": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsAborted(); !ok {
		return &ValidationError{Name: "is_aborted", err: errors.New(`ent: missing required field "Conversations.is_aborted"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Conversations.created_at"`)}
	}
	return nil
}

func (cc *ConversationsCreate) sqlSave(ctx context.Context) (*Conversations, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConversationsCreate) createSpec() (*Conversations, *sqlgraph.CreateSpec) {
	var (
		_node = &Conversations{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(conversations.Table, sqlgraph.NewFieldSpec(conversations.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.AiModel(); ok {
		_spec.SetField(conversations.FieldAiModel, field.TypeEnum, value)
		_node.AiModel = value
	}
	if value, ok := cc.mutation.SnsType(); ok {
		_spec.SetField(conversations.FieldSnsType, field.TypeEnum, value)
		_node.SnsType = value
	}
	if value, ok := cc.mutation.CmdVersion(); ok {
		_spec.SetField(conversations.FieldCmdVersion, field.TypeEnum, value)
		_node.CmdVersion = value
	}
	if value, ok := cc.mutation.IsAborted(); ok {
		_spec.SetField(conversations.FieldIsAborted, field.TypeBool, value)
		_node.IsAborted = value
	}
	if value, ok := cc.mutation.AbortReason(); ok {
		_spec.SetField(conversations.FieldAbortReason, field.TypeString, value)
		_node.AbortReason = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(conversations.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ConversationsCreateBulk is the builder for creating many Conversations entities in bulk.
type ConversationsCreateBulk struct {
	config
	builders []*ConversationsCreate
}

// Save creates the Conversations entities in the database.
func (ccb *ConversationsCreateBulk) Save(ctx context.Context) ([]*Conversations, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Conversations, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConversationsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConversationsCreateBulk) SaveX(ctx context.Context) []*Conversations {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConversationsCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConversationsCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}