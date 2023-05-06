// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/chatgpt35turboconversationlog"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/conversations"
)

// Chatgpt35TurboConversationLog is the model entity for the Chatgpt35TurboConversationLog schema.
type Chatgpt35TurboConversationLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Purpose holds the value of the "purpose" field.
	Purpose string `json:"purpose,omitempty"`
	// Role holds the value of the "role" field.
	Role chatgpt35turboconversationlog.Role `json:"role,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Chatgpt35TurboConversationLogQuery when eager-loading is set.
	Edges                                        Chatgpt35TurboConversationLogEdges `json:"edges"`
	chatgpt35turbo_conversation_log_conversation *int
	selectValues                                 sql.SelectValues
}

// Chatgpt35TurboConversationLogEdges holds the relations/edges for other nodes in the graph.
type Chatgpt35TurboConversationLogEdges struct {
	// Conversation holds the value of the conversation edge.
	Conversation *Conversations `json:"conversation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ConversationOrErr returns the Conversation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Chatgpt35TurboConversationLogEdges) ConversationOrErr() (*Conversations, error) {
	if e.loadedTypes[0] {
		if e.Conversation == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: conversations.Label}
		}
		return e.Conversation, nil
	}
	return nil, &NotLoadedError{edge: "conversation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chatgpt35TurboConversationLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chatgpt35turboconversationlog.FieldID:
			values[i] = new(sql.NullInt64)
		case chatgpt35turboconversationlog.FieldMessage, chatgpt35turboconversationlog.FieldPurpose, chatgpt35turboconversationlog.FieldRole:
			values[i] = new(sql.NullString)
		case chatgpt35turboconversationlog.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case chatgpt35turboconversationlog.ForeignKeys[0]: // chatgpt35turbo_conversation_log_conversation
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chatgpt35TurboConversationLog fields.
func (ccl *Chatgpt35TurboConversationLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chatgpt35turboconversationlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ccl.ID = int(value.Int64)
		case chatgpt35turboconversationlog.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				ccl.Message = value.String
			}
		case chatgpt35turboconversationlog.FieldPurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose", values[i])
			} else if value.Valid {
				ccl.Purpose = value.String
			}
		case chatgpt35turboconversationlog.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				ccl.Role = chatgpt35turboconversationlog.Role(value.String)
			}
		case chatgpt35turboconversationlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ccl.CreatedAt = value.Time
			}
		case chatgpt35turboconversationlog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chatgpt35turbo_conversation_log_conversation", value)
			} else if value.Valid {
				ccl.chatgpt35turbo_conversation_log_conversation = new(int)
				*ccl.chatgpt35turbo_conversation_log_conversation = int(value.Int64)
			}
		default:
			ccl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Chatgpt35TurboConversationLog.
// This includes values selected through modifiers, order, etc.
func (ccl *Chatgpt35TurboConversationLog) Value(name string) (ent.Value, error) {
	return ccl.selectValues.Get(name)
}

// QueryConversation queries the "conversation" edge of the Chatgpt35TurboConversationLog entity.
func (ccl *Chatgpt35TurboConversationLog) QueryConversation() *ConversationsQuery {
	return NewChatgpt35TurboConversationLogClient(ccl.config).QueryConversation(ccl)
}

// Update returns a builder for updating this Chatgpt35TurboConversationLog.
// Note that you need to call Chatgpt35TurboConversationLog.Unwrap() before calling this method if this Chatgpt35TurboConversationLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ccl *Chatgpt35TurboConversationLog) Update() *Chatgpt35TurboConversationLogUpdateOne {
	return NewChatgpt35TurboConversationLogClient(ccl.config).UpdateOne(ccl)
}

// Unwrap unwraps the Chatgpt35TurboConversationLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ccl *Chatgpt35TurboConversationLog) Unwrap() *Chatgpt35TurboConversationLog {
	_tx, ok := ccl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chatgpt35TurboConversationLog is not a transactional entity")
	}
	ccl.config.driver = _tx.drv
	return ccl
}

// String implements the fmt.Stringer.
func (ccl *Chatgpt35TurboConversationLog) String() string {
	var builder strings.Builder
	builder.WriteString("Chatgpt35TurboConversationLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ccl.ID))
	builder.WriteString("message=")
	builder.WriteString(ccl.Message)
	builder.WriteString(", ")
	builder.WriteString("purpose=")
	builder.WriteString(ccl.Purpose)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", ccl.Role))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ccl.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Chatgpt35TurboConversationLogs is a parsable slice of Chatgpt35TurboConversationLog.
type Chatgpt35TurboConversationLogs []*Chatgpt35TurboConversationLog
