// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/ent/chatgpt35turboconversationlog"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/ent/conversations"
)

// Chatgpt35TurboConversationLogCreate is the builder for creating a Chatgpt35TurboConversationLog entity.
type Chatgpt35TurboConversationLogCreate struct {
	config
	mutation *Chatgpt35TurboConversationLogMutation
	hooks    []Hook
}

// SetMessage sets the "message" field.
func (cclc *Chatgpt35TurboConversationLogCreate) SetMessage(s string) *Chatgpt35TurboConversationLogCreate {
	cclc.mutation.SetMessage(s)
	return cclc
}

// SetPurpose sets the "purpose" field.
func (cclc *Chatgpt35TurboConversationLogCreate) SetPurpose(s string) *Chatgpt35TurboConversationLogCreate {
	cclc.mutation.SetPurpose(s)
	return cclc
}

// SetNillablePurpose sets the "purpose" field if the given value is not nil.
func (cclc *Chatgpt35TurboConversationLogCreate) SetNillablePurpose(s *string) *Chatgpt35TurboConversationLogCreate {
	if s != nil {
		cclc.SetPurpose(*s)
	}
	return cclc
}

// SetRole sets the "role" field.
func (cclc *Chatgpt35TurboConversationLogCreate) SetRole(c chatgpt35turboconversationlog.Role) *Chatgpt35TurboConversationLogCreate {
	cclc.mutation.SetRole(c)
	return cclc
}

// SetCreatedAt sets the "created_at" field.
func (cclc *Chatgpt35TurboConversationLogCreate) SetCreatedAt(t time.Time) *Chatgpt35TurboConversationLogCreate {
	cclc.mutation.SetCreatedAt(t)
	return cclc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cclc *Chatgpt35TurboConversationLogCreate) SetNillableCreatedAt(t *time.Time) *Chatgpt35TurboConversationLogCreate {
	if t != nil {
		cclc.SetCreatedAt(*t)
	}
	return cclc
}

// SetConversationID sets the "conversation" edge to the Conversations entity by ID.
func (cclc *Chatgpt35TurboConversationLogCreate) SetConversationID(id int) *Chatgpt35TurboConversationLogCreate {
	cclc.mutation.SetConversationID(id)
	return cclc
}

// SetNillableConversationID sets the "conversation" edge to the Conversations entity by ID if the given value is not nil.
func (cclc *Chatgpt35TurboConversationLogCreate) SetNillableConversationID(id *int) *Chatgpt35TurboConversationLogCreate {
	if id != nil {
		cclc = cclc.SetConversationID(*id)
	}
	return cclc
}

// SetConversation sets the "conversation" edge to the Conversations entity.
func (cclc *Chatgpt35TurboConversationLogCreate) SetConversation(c *Conversations) *Chatgpt35TurboConversationLogCreate {
	return cclc.SetConversationID(c.ID)
}

// Mutation returns the Chatgpt35TurboConversationLogMutation object of the builder.
func (cclc *Chatgpt35TurboConversationLogCreate) Mutation() *Chatgpt35TurboConversationLogMutation {
	return cclc.mutation
}

// Save creates the Chatgpt35TurboConversationLog in the database.
func (cclc *Chatgpt35TurboConversationLogCreate) Save(ctx context.Context) (*Chatgpt35TurboConversationLog, error) {
	cclc.defaults()
	return withHooks(ctx, cclc.sqlSave, cclc.mutation, cclc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cclc *Chatgpt35TurboConversationLogCreate) SaveX(ctx context.Context) *Chatgpt35TurboConversationLog {
	v, err := cclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cclc *Chatgpt35TurboConversationLogCreate) Exec(ctx context.Context) error {
	_, err := cclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cclc *Chatgpt35TurboConversationLogCreate) ExecX(ctx context.Context) {
	if err := cclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cclc *Chatgpt35TurboConversationLogCreate) defaults() {
	if _, ok := cclc.mutation.CreatedAt(); !ok {
		v := chatgpt35turboconversationlog.DefaultCreatedAt()
		cclc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cclc *Chatgpt35TurboConversationLogCreate) check() error {
	if _, ok := cclc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "Chatgpt35TurboConversationLog.message"`)}
	}
	if v, ok := cclc.mutation.Message(); ok {
		if err := chatgpt35turboconversationlog.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Chatgpt35TurboConversationLog.message": %w`, err)}
		}
	}
	if _, ok := cclc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Chatgpt35TurboConversationLog.role"`)}
	}
	if v, ok := cclc.mutation.Role(); ok {
		if err := chatgpt35turboconversationlog.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Chatgpt35TurboConversationLog.role": %w`, err)}
		}
	}
	if _, ok := cclc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Chatgpt35TurboConversationLog.created_at"`)}
	}
	return nil
}

func (cclc *Chatgpt35TurboConversationLogCreate) sqlSave(ctx context.Context) (*Chatgpt35TurboConversationLog, error) {
	if err := cclc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cclc.mutation.id = &_node.ID
	cclc.mutation.done = true
	return _node, nil
}

func (cclc *Chatgpt35TurboConversationLogCreate) createSpec() (*Chatgpt35TurboConversationLog, *sqlgraph.CreateSpec) {
	var (
		_node = &Chatgpt35TurboConversationLog{config: cclc.config}
		_spec = sqlgraph.NewCreateSpec(chatgpt35turboconversationlog.Table, sqlgraph.NewFieldSpec(chatgpt35turboconversationlog.FieldID, field.TypeInt))
	)
	if value, ok := cclc.mutation.Message(); ok {
		_spec.SetField(chatgpt35turboconversationlog.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := cclc.mutation.Purpose(); ok {
		_spec.SetField(chatgpt35turboconversationlog.FieldPurpose, field.TypeString, value)
		_node.Purpose = value
	}
	if value, ok := cclc.mutation.Role(); ok {
		_spec.SetField(chatgpt35turboconversationlog.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := cclc.mutation.CreatedAt(); ok {
		_spec.SetField(chatgpt35turboconversationlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := cclc.mutation.ConversationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   chatgpt35turboconversationlog.ConversationTable,
			Columns: []string{chatgpt35turboconversationlog.ConversationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(conversations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.chatgpt35turbo_conversation_log_conversation = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Chatgpt35TurboConversationLogCreateBulk is the builder for creating many Chatgpt35TurboConversationLog entities in bulk.
type Chatgpt35TurboConversationLogCreateBulk struct {
	config
	builders []*Chatgpt35TurboConversationLogCreate
}

// Save creates the Chatgpt35TurboConversationLog entities in the database.
func (cclcb *Chatgpt35TurboConversationLogCreateBulk) Save(ctx context.Context) ([]*Chatgpt35TurboConversationLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cclcb.builders))
	nodes := make([]*Chatgpt35TurboConversationLog, len(cclcb.builders))
	mutators := make([]Mutator, len(cclcb.builders))
	for i := range cclcb.builders {
		func(i int, root context.Context) {
			builder := cclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Chatgpt35TurboConversationLogMutation)
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
					_, err = mutators[i+1].Mutate(root, cclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cclcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cclcb *Chatgpt35TurboConversationLogCreateBulk) SaveX(ctx context.Context) []*Chatgpt35TurboConversationLog {
	v, err := cclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cclcb *Chatgpt35TurboConversationLogCreateBulk) Exec(ctx context.Context) error {
	_, err := cclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cclcb *Chatgpt35TurboConversationLogCreateBulk) ExecX(ctx context.Context) {
	if err := cclcb.Exec(ctx); err != nil {
		panic(err)
	}
}