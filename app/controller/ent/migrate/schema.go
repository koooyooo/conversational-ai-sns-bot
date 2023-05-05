// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// Chatgpt35turboConversationLogsColumns holds the columns for the "chatgpt35turbo_conversation_logs" table.
	Chatgpt35turboConversationLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message", Type: field.TypeString},
		{Name: "purpose", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"system", "user", "assistant"}},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "chatgpt35turbo_conversation_log_conversation", Type: field.TypeInt, Nullable: true},
	}
	// Chatgpt35turboConversationLogsTable holds the schema information for the "chatgpt35turbo_conversation_logs" table.
	Chatgpt35turboConversationLogsTable = &schema.Table{
		Name:       "chatgpt35turbo_conversation_logs",
		Columns:    Chatgpt35turboConversationLogsColumns,
		PrimaryKey: []*schema.Column{Chatgpt35turboConversationLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chatgpt35turbo_conversation_logs_conversations_conversation",
				Columns:    []*schema.Column{Chatgpt35turboConversationLogsColumns[5]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConversationsColumns holds the columns for the "conversations" table.
	ConversationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ai_model", Type: field.TypeEnum, Enums: []string{"gpt-3_5-turbo"}},
		{Name: "sns_type", Type: field.TypeEnum, Enums: []string{"twitter"}},
		{Name: "cmd_version", Type: field.TypeEnum, Enums: []string{"v0_1"}},
		{Name: "is_aborted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// ConversationsTable holds the schema information for the "conversations" table.
	ConversationsTable = &schema.Table{
		Name:       "conversations",
		Columns:    ConversationsColumns,
		PrimaryKey: []*schema.Column{ConversationsColumns[0]},
	}
	// TwitterAccountsColumns holds the columns for the "twitter_accounts" table.
	TwitterAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "twitter_account_id", Type: field.TypeString, Unique: true},
		{Name: "bearer_token", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "twitter_accounts_conversation", Type: field.TypeInt, Nullable: true},
	}
	// TwitterAccountsTable holds the schema information for the "twitter_accounts" table.
	TwitterAccountsTable = &schema.Table{
		Name:       "twitter_accounts",
		Columns:    TwitterAccountsColumns,
		PrimaryKey: []*schema.Column{TwitterAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "twitter_accounts_conversations_conversation",
				Columns:    []*schema.Column{TwitterAccountsColumns[5]},
				RefColumns: []*schema.Column{ConversationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		Chatgpt35turboConversationLogsTable,
		ConversationsTable,
		TwitterAccountsTable,
	}
)

func init() {
	Chatgpt35turboConversationLogsTable.ForeignKeys[0].RefTable = ConversationsTable
	TwitterAccountsTable.ForeignKeys[0].RefTable = ConversationsTable
}
