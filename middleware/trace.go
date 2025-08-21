package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zsljava/gokit/constant"
	"github.com/zsljava/gokit/log"
	"go.uber.org/zap"
)

// GinTraceMiddleware Gin  traceId 中间件
func GinTraceMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 traceId（如果存在）
		traceID := c.GetHeader(constant.XTraceId)
		if traceID == "" {
			// 生成新的 traceId
			traceID = uuid.New().String()
		}

		// 将 traceId 设置到 Gin context 和标准 context
		c.Set(constant.TraceIDKey, traceID)

		// 设置响应头
		c.Header(constant.XTraceId, traceID)

		logger.WithValue(c, zap.String(constant.TraceIDKey, traceID))
		c.Next()
	}
}
