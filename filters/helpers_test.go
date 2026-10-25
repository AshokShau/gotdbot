package filters_test

import (
	"reflect"
	"testing"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters"
)

func TestFilterCompositionShortCircuitAndSnapshot(t *testing.T) {
	var calls []int
	first := func(int) bool { calls = append(calls, 1); return false }
	second := func(int) bool { calls = append(calls, 2); return true }
	if filters.AllFilters(first, second)(0) || !reflect.DeepEqual(calls, []int{1}) {
		t.Fatal("AND did not short-circuit")
	}
	calls = nil
	if !filters.AnyFilters(second, first)(0) || !reflect.DeepEqual(calls, []int{2}) {
		t.Fatal("OR did not short-circuit")
	}
	if !filters.AllFilters[int](nil)(0) || filters.AnyFilters[int](nil)(0) || filters.NotFilter[int](nil)(0) {
		t.Fatal("nil filter semantics changed")
	}
	predicates := []func(int) bool{second}
	all, any := filters.AllFilters(predicates...), filters.AnyFilters(predicates...)
	predicates[0] = first
	if !all(0) || !any(0) {
		t.Fatal("composition retained caller's mutable slice")
	}
	var private gotdbot.MessageFilter = func(m *gotdbot.Message) bool { return m.IsPrivate() }
	var incoming gotdbot.MessageFilter = func(m *gotdbot.Message) bool { return !m.IsOutgoingMessage() }
	if !private.And(incoming)(&gotdbot.Message{ChatId: 12}) || private.Or(incoming.Not())(&gotdbot.Message{ChatId: -12}) {
		t.Fatal("message filter methods composed incorrectly")
	}
	var callback gotdbot.UpdateNewCallbackQueryFilter = func(u *gotdbot.UpdateNewCallbackQuery) bool { return u.SenderUserId == 7 }
	// The generic helpers work with existing named filters without conversions.
	combined := filters.AllFilters(callback, func(u *gotdbot.UpdateNewCallbackQuery) bool { return u.ChatId == 12 })
	if !combined(&gotdbot.UpdateNewCallbackQuery{ChatId: 12, SenderUserId: 7}) || combined(&gotdbot.UpdateNewCallbackQuery{ChatId: 13, SenderUserId: 7}) {
		t.Fatal("typed update filters composed incorrectly")
	}
}
