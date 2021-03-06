// Code generated by "go-syncmap -output context_sync_map.gen.go -type contextMap<string,context.Context>"; DO NOT EDIT.

package expire

import (
	"context"
	"sync"
)

func _() {
	// An "cannot convert contextMap literal (type contextMap) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)(contextMap{})
}

var _nil_contextMap_context_Context_value = func() (val context.Context) { return }()

func (m *contextMap) Store(key string, value context.Context) {
	(*sync.Map)(m).Store(key, value)
}

func (m *contextMap) LoadOrStore(key string, value context.Context) (context.Context, bool) {
	actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
		return _nil_contextMap_context_Context_value, loaded
	}
	return actual.(context.Context), loaded
}

func (m *contextMap) Load(key string) (context.Context, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if value == nil {
		return _nil_contextMap_context_Context_value, ok
	}
	return value.(context.Context), ok
}

func (m *contextMap) Delete(key string) {
	(*sync.Map)(m).Delete(key)
}

func (m *contextMap) Range(f func(key string, value context.Context) bool) {
	(*sync.Map)(m).Range(func(key, value interface{}) bool {
		return f(key.(string), value.(context.Context))
	})
}
