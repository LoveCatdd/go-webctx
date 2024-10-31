package goroutine

import "strings"

const (
	JWT_MAP_CLAIM             = "map_claims"
	IDENTITY_CONTEXT_INFO_KEY = "identity_context_info"
)

type GoroutineContextHolder struct {
	_map *safemap
}

func (g *GoroutineContextHolder) Initialization() {
	g._map = newSafemap()
}

func (g *GoroutineContextHolder) ContenxtMap() *safemap {
	return g._map
}

func (g *GoroutineContextHolder) SetContextMap(m *safemap) {
	if m == nil {
		g._map = newSafemap()
	} else {
		g._map = m
	}
}

func (g *GoroutineContextHolder) ContextWithKey(key string) (any, bool) {
	return g.ContenxtMap().Get(key)
}

func (g *GoroutineContextHolder) Change(key string, cInfo any) string {
	if strings.TrimSpace(key) == "" || cInfo == nil {
		g.ContenxtMap().Remove(key)
		return "change remove"
	}
	g.ContenxtMap().Set(key, cInfo)
	return "change set"
}

func (g *GoroutineContextHolder) Remove(key string) {
	g.ContenxtMap().Remove(key)
}

func (g *GoroutineContextHolder) Clear() {
	g.ContenxtMap().Clear()
}
