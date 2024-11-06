package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/LoveCatdd/util/pkg/lib/core/ids"
	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"github.com/gin-gonic/gin"
)

func TraceMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()
		log.SetTraceId(ids.UUIDV1())

		// 创建自定义的 ResponseWriter
		customWriter := &customResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		// 替换 Gin 的默认 ResponseWriter
		c.Writer = customWriter

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		request_body, _ := io.ReadAll(c.Request.Body)
		log.
			Info("requst_info start: path: %v, method: %v, ip: %v, query: %v, request_body: %v",
				path, c.Request.Method,
				c.ClientIP(), query, string(request_body))

			// 重新获取
		c.Request.Body = io.NopCloser(bytes.NewReader(request_body))

		c.Next()

		cost := time.Since(start)
		log.
			Info("requst_info end: status: %v, path: %v, method: %v, ip: %v, cost: %v, response_body: %v",
				c.Writer.Status(), path, c.Request.Method,
				c.ClientIP(), cost,
				customWriter.body.String())
	}
}

// 自定义一个 ResponseWriter 用于捕获响应内容
type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 写入响应内容
func (w *customResponseWriter) Write(p []byte) (n int, err error) {
	// 将响应内容写入自定义的 body 缓冲区
	w.body.Write(p)
	return w.ResponseWriter.Write(p)
}
