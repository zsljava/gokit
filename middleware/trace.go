package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zsljava/gokit/global"
	"github.com/zsljava/gokit/util/log"
	"go.uber.org/zap"
)

// GinTraceMiddleware Gin  traceId 中间件
func GinTraceMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 traceId（如果存在）
		traceID := c.GetHeader(global.XTraceId)
		if traceID == "" {
			// 生成新的 traceId
			traceID = uuid.New().String()
		}

		// 将 traceId 设置到 Gin context 和标准 context
		c.Set(global.TraceIDKey, traceID)

		// 设置响应头
		c.Header(global.XTraceId, traceID)

		logger.WithValue(c, zap.String(global.TraceIDKey, traceID))
		c.Next()
	}
}
