package context

import "github.com/LoveCatdd/webctx/pkg/lib/core/goroutine"

const (
	CustonContextKey ContextValueKey = "CustonContext"
)

// 自定义 context
type CustomContext struct {

	// 协程上下文
	contextHolder *goroutine.GoroutineContextHolder
}

// 生成实例
func NewCustomContext(contextHolder *goroutine.GoroutineContextHolder) *CustomContext {

	return &CustomContext{
		contextHolder: contextHolder,
	}
}

func (c *CustomContext) ContextHolder() *goroutine.GoroutineContextHolder {
	return c.contextHolder
}
