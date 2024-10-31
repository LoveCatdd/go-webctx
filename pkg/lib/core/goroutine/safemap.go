package goroutine

import (
	"sync"
)

type safemap struct {
	m sync.Map
}

func newSafemap() *safemap {
	return &safemap{}
}

func (s *safemap) Get(key string) (any, bool) {
	return s.m.Load(key)
}

func (s *safemap) Remove(key string) {
	s.m.Delete(key)
}

func (s *safemap) Set(key string, value any) {
	s.m.Store(key, value)
}

func (s *safemap) Clear() {
	s.m.Range(func(key, value any) bool {
		s.Remove(key.(string))
		return true
	})
}
