// Code generated by ent, DO NOT EDIT.

package conversations

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the conversations type in the database.
	Label = "conversations"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAiModel holds the string denoting the ai_model field in the database.
	FieldAiModel = "ai_model"
	// FieldSnsType holds the string denoting the sns_type field in the database.
	FieldSnsType = "sns_type"
	// FieldCmdVersion holds the string denoting the cmd_version field in the database.
	FieldCmdVersion = "cmd_version"
	// FieldIsAborted holds the string denoting the is_aborted field in the database.
	FieldIsAborted = "is_aborted"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the conversations in the database.
	Table = "conversations"
)

// Columns holds all SQL columns for conversations fields.
var Columns = []string{
	FieldID,
	FieldAiModel,
	FieldSnsType,
	FieldCmdVersion,
	FieldIsAborted,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsAborted holds the default value on creation for the "is_aborted" field.
	DefaultIsAborted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// AiModel defines the type for the "ai_model" enum field.
type AiModel string

// AiModel values.
const (
	AiModelGpt35Turbo AiModel = "gpt-3_5-turbo"
)

func (am AiModel) String() string {
	return string(am)
}

// AiModelValidator is a validator for the "ai_model" field enum values. It is called by the builders before save.
func AiModelValidator(am AiModel) error {
	switch am {
	case AiModelGpt35Turbo:
		return nil
	default:
		return fmt.Errorf("conversations: invalid enum value for ai_model field: %q", am)
	}
}

// SnsType defines the type for the "sns_type" enum field.
type SnsType string

// SnsType values.
const (
	SnsTypeTwitter SnsType = "twitter"
)

func (st SnsType) String() string {
	return string(st)
}

// SnsTypeValidator is a validator for the "sns_type" field enum values. It is called by the builders before save.
func SnsTypeValidator(st SnsType) error {
	switch st {
	case SnsTypeTwitter:
		return nil
	default:
		return fmt.Errorf("conversations: invalid enum value for sns_type field: %q", st)
	}
}

// CmdVersion defines the type for the "cmd_version" enum field.
type CmdVersion string

// CmdVersion values.
const (
	CmdVersionV01 CmdVersion = "v0_1"
)

func (cv CmdVersion) String() string {
	return string(cv)
}

// CmdVersionValidator is a validator for the "cmd_version" field enum values. It is called by the builders before save.
func CmdVersionValidator(cv CmdVersion) error {
	switch cv {
	case CmdVersionV01:
		return nil
	default:
		return fmt.Errorf("conversations: invalid enum value for cmd_version field: %q", cv)
	}
}

// OrderOption defines the ordering options for the Conversations queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAiModel orders the results by the ai_model field.
func ByAiModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAiModel, opts...).ToFunc()
}

// BySnsType orders the results by the sns_type field.
func BySnsType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnsType, opts...).ToFunc()
}

// ByCmdVersion orders the results by the cmd_version field.
func ByCmdVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCmdVersion, opts...).ToFunc()
}

// ByIsAborted orders the results by the is_aborted field.
func ByIsAborted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAborted, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
