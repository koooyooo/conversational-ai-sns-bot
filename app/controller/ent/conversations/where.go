// Code generated by ent, DO NOT EDIT.

package conversations

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/YukiTsuchida/conversational-ai-sns-bot/app/controller/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Conversations {
	return predicate.Conversations(sql.FieldLTE(FieldID, id))
}

// IsAborted applies equality check predicate on the "is_aborted" field. It's identical to IsAbortedEQ.
func IsAborted(v bool) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldIsAborted, v))
}

// AbortReason applies equality check predicate on the "abort_reason" field. It's identical to AbortReasonEQ.
func AbortReason(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldAbortReason, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldCreatedAt, v))
}

// AiModelEQ applies the EQ predicate on the "ai_model" field.
func AiModelEQ(v AiModel) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldAiModel, v))
}

// AiModelNEQ applies the NEQ predicate on the "ai_model" field.
func AiModelNEQ(v AiModel) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldAiModel, v))
}

// AiModelIn applies the In predicate on the "ai_model" field.
func AiModelIn(vs ...AiModel) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldAiModel, vs...))
}

// AiModelNotIn applies the NotIn predicate on the "ai_model" field.
func AiModelNotIn(vs ...AiModel) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldAiModel, vs...))
}

// SnsTypeEQ applies the EQ predicate on the "sns_type" field.
func SnsTypeEQ(v SnsType) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldSnsType, v))
}

// SnsTypeNEQ applies the NEQ predicate on the "sns_type" field.
func SnsTypeNEQ(v SnsType) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldSnsType, v))
}

// SnsTypeIn applies the In predicate on the "sns_type" field.
func SnsTypeIn(vs ...SnsType) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldSnsType, vs...))
}

// SnsTypeNotIn applies the NotIn predicate on the "sns_type" field.
func SnsTypeNotIn(vs ...SnsType) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldSnsType, vs...))
}

// CmdVersionEQ applies the EQ predicate on the "cmd_version" field.
func CmdVersionEQ(v CmdVersion) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldCmdVersion, v))
}

// CmdVersionNEQ applies the NEQ predicate on the "cmd_version" field.
func CmdVersionNEQ(v CmdVersion) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldCmdVersion, v))
}

// CmdVersionIn applies the In predicate on the "cmd_version" field.
func CmdVersionIn(vs ...CmdVersion) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldCmdVersion, vs...))
}

// CmdVersionNotIn applies the NotIn predicate on the "cmd_version" field.
func CmdVersionNotIn(vs ...CmdVersion) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldCmdVersion, vs...))
}

// IsAbortedEQ applies the EQ predicate on the "is_aborted" field.
func IsAbortedEQ(v bool) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldIsAborted, v))
}

// IsAbortedNEQ applies the NEQ predicate on the "is_aborted" field.
func IsAbortedNEQ(v bool) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldIsAborted, v))
}

// AbortReasonEQ applies the EQ predicate on the "abort_reason" field.
func AbortReasonEQ(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldAbortReason, v))
}

// AbortReasonNEQ applies the NEQ predicate on the "abort_reason" field.
func AbortReasonNEQ(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldAbortReason, v))
}

// AbortReasonIn applies the In predicate on the "abort_reason" field.
func AbortReasonIn(vs ...string) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldAbortReason, vs...))
}

// AbortReasonNotIn applies the NotIn predicate on the "abort_reason" field.
func AbortReasonNotIn(vs ...string) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldAbortReason, vs...))
}

// AbortReasonGT applies the GT predicate on the "abort_reason" field.
func AbortReasonGT(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldGT(FieldAbortReason, v))
}

// AbortReasonGTE applies the GTE predicate on the "abort_reason" field.
func AbortReasonGTE(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldGTE(FieldAbortReason, v))
}

// AbortReasonLT applies the LT predicate on the "abort_reason" field.
func AbortReasonLT(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldLT(FieldAbortReason, v))
}

// AbortReasonLTE applies the LTE predicate on the "abort_reason" field.
func AbortReasonLTE(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldLTE(FieldAbortReason, v))
}

// AbortReasonContains applies the Contains predicate on the "abort_reason" field.
func AbortReasonContains(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldContains(FieldAbortReason, v))
}

// AbortReasonHasPrefix applies the HasPrefix predicate on the "abort_reason" field.
func AbortReasonHasPrefix(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldHasPrefix(FieldAbortReason, v))
}

// AbortReasonHasSuffix applies the HasSuffix predicate on the "abort_reason" field.
func AbortReasonHasSuffix(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldHasSuffix(FieldAbortReason, v))
}

// AbortReasonIsNil applies the IsNil predicate on the "abort_reason" field.
func AbortReasonIsNil() predicate.Conversations {
	return predicate.Conversations(sql.FieldIsNull(FieldAbortReason))
}

// AbortReasonNotNil applies the NotNil predicate on the "abort_reason" field.
func AbortReasonNotNil() predicate.Conversations {
	return predicate.Conversations(sql.FieldNotNull(FieldAbortReason))
}

// AbortReasonEqualFold applies the EqualFold predicate on the "abort_reason" field.
func AbortReasonEqualFold(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldEqualFold(FieldAbortReason, v))
}

// AbortReasonContainsFold applies the ContainsFold predicate on the "abort_reason" field.
func AbortReasonContainsFold(v string) predicate.Conversations {
	return predicate.Conversations(sql.FieldContainsFold(FieldAbortReason, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Conversations {
	return predicate.Conversations(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Conversations) predicate.Conversations {
	return predicate.Conversations(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Conversations) predicate.Conversations {
	return predicate.Conversations(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Conversations) predicate.Conversations {
	return predicate.Conversations(func(s *sql.Selector) {
		p(s.Not())
	})
}
