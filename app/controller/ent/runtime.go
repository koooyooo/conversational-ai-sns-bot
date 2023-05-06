// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/chatgpt35turboconversationlog"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/conversations"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/schema"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/twitteraccounts"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatgpt35turboconversationlogFields := schema.Chatgpt35TurboConversationLog{}.Fields()
	_ = chatgpt35turboconversationlogFields
	// chatgpt35turboconversationlogDescMessage is the schema descriptor for message field.
	chatgpt35turboconversationlogDescMessage := chatgpt35turboconversationlogFields[0].Descriptor()
	// chatgpt35turboconversationlog.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	chatgpt35turboconversationlog.MessageValidator = chatgpt35turboconversationlogDescMessage.Validators[0].(func(string) error)
	// chatgpt35turboconversationlogDescCreatedAt is the schema descriptor for created_at field.
	chatgpt35turboconversationlogDescCreatedAt := chatgpt35turboconversationlogFields[3].Descriptor()
	// chatgpt35turboconversationlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	chatgpt35turboconversationlog.DefaultCreatedAt = chatgpt35turboconversationlogDescCreatedAt.Default.(func() time.Time)
	conversationsFields := schema.Conversations{}.Fields()
	_ = conversationsFields
	// conversationsDescIsAborted is the schema descriptor for is_aborted field.
	conversationsDescIsAborted := conversationsFields[3].Descriptor()
	// conversations.DefaultIsAborted holds the default value on creation for the is_aborted field.
	conversations.DefaultIsAborted = conversationsDescIsAborted.Default.(bool)
	// conversationsDescCreatedAt is the schema descriptor for created_at field.
	conversationsDescCreatedAt := conversationsFields[4].Descriptor()
	// conversations.DefaultCreatedAt holds the default value on creation for the created_at field.
	conversations.DefaultCreatedAt = conversationsDescCreatedAt.Default.(func() time.Time)
	twitteraccountsFields := schema.TwitterAccounts{}.Fields()
	_ = twitteraccountsFields
	// twitteraccountsDescTwitterAccountID is the schema descriptor for twitter_account_id field.
	twitteraccountsDescTwitterAccountID := twitteraccountsFields[0].Descriptor()
	// twitteraccounts.TwitterAccountIDValidator is a validator for the "twitter_account_id" field. It is called by the builders before save.
	twitteraccounts.TwitterAccountIDValidator = twitteraccountsDescTwitterAccountID.Validators[0].(func(string) error)
	// twitteraccountsDescBearerToken is the schema descriptor for bearer_token field.
	twitteraccountsDescBearerToken := twitteraccountsFields[1].Descriptor()
	// twitteraccounts.BearerTokenValidator is a validator for the "bearer_token" field. It is called by the builders before save.
	twitteraccounts.BearerTokenValidator = twitteraccountsDescBearerToken.Validators[0].(func(string) error)
	// twitteraccountsDescCreatedAt is the schema descriptor for created_at field.
	twitteraccountsDescCreatedAt := twitteraccountsFields[2].Descriptor()
	// twitteraccounts.DefaultCreatedAt holds the default value on creation for the created_at field.
	twitteraccounts.DefaultCreatedAt = twitteraccountsDescCreatedAt.Default.(func() time.Time)
	// twitteraccountsDescUpdatedAt is the schema descriptor for updated_at field.
	twitteraccountsDescUpdatedAt := twitteraccountsFields[3].Descriptor()
	// twitteraccounts.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	twitteraccounts.DefaultUpdatedAt = twitteraccountsDescUpdatedAt.Default.(func() time.Time)
}
