// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ConversationsColumns holds the columns for the "conversations" table.
	ConversationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ai_model", Type: field.TypeEnum, Enums: []string{"chatgpt-3_5-turbo"}},
		{Name: "sns_type", Type: field.TypeEnum, Enums: []string{"twitter"}},
		{Name: "cmd_version", Type: field.TypeEnum, Enums: []string{"v0_1"}},
		{Name: "is_aborted", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "twitter_accounts_conversation", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ConversationsTable holds the schema information for the "conversations" table.
	ConversationsTable = &schema.Table{
		Name:       "conversations",
		Columns:    ConversationsColumns,
		PrimaryKey: []*schema.Column{ConversationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "conversations_twitter_accounts_conversation",
				Columns:    []*schema.Column{ConversationsColumns[6]},
				RefColumns: []*schema.Column{TwitterAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TwitterAccountsColumns holds the columns for the "twitter_accounts" table.
	TwitterAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "twitter_account_id", Type: field.TypeString, Unique: true},
		{Name: "bearer_token", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// TwitterAccountsTable holds the schema information for the "twitter_accounts" table.
	TwitterAccountsTable = &schema.Table{
		Name:       "twitter_accounts",
		Columns:    TwitterAccountsColumns,
		PrimaryKey: []*schema.Column{TwitterAccountsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConversationsTable,
		TwitterAccountsTable,
	}
)

func init() {
	ConversationsTable.ForeignKeys[0].RefTable = TwitterAccountsTable
}