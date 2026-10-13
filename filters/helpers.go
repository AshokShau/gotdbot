// Package filters composes predicates for messages, typed updates, and other values.
// Message and callback predicates are available in the message and callbackquery
// subpackages.
package filters

// AllFilters matches when all non-nil filters match. It short-circuits on the
// first false result. An empty set of filters matches everything.
func AllFilters[T any](filters ...func(T) bool) func(T) bool {
	predicates := append([]func(T) bool(nil), filters...)
	return func(value T) bool {
		for _, filter := range predicates {
			if filter != nil && !filter(value) {
				return false
			}
		}
		return true
	}
}

// AnyFilters matches when any non-nil filter matches. It short-circuits on the
// first true result. An empty set of filters matches nothing.
func AnyFilters[T any](filters ...func(T) bool) func(T) bool {
	predicates := append([]func(T) bool(nil), filters...)
	return func(value T) bool {
		for _, filter := range predicates {
			if filter != nil && filter(value) {
				return true
			}
		}
		return false
	}
}

// NotFilter negates a filter. A nil filter matches nothing.
func NotFilter[T any](filter func(T) bool) func(T) bool {
	return func(value T) bool { return filter != nil && !filter(value) }
}
