package supressor

import (
  "sync"
)

// SeenFilter is a simple data structure to check whether
// certain keys have been seen before. Useful to implement
// idempotent apis, and supressor filters.
type SeenFilter interface {

  // Filter adds key k to the filter, and
  // returns true if this key had been seen before.
  // (false means the key was seen for the first time)
  Filter(k string) (new bool)
}

// make faster with bloom filters or whatever.
// this is a reference impl, and needs optimization.
type seenFilter struct {
  M map[string]struct{}

  sync.Mutex
}

func NewSeenFilter() *seenFilter {
  return &seenFilter{M: make(map[string]struct{})}
}

func (f *seenFilter) Filter(k string) (new bool) {
  f.Lock()
  _, found := f.M[k]
  if !found {
    f.M[k] = struct{}{}
  }
  f.Unlock()
  return !found
}
