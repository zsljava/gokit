package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/zsljava/gokit/util/log"
	"go.uber.org/zap"
	"io"
	"time"
)

func RequestLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ""
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			body = string(bodyBytes)
		}

		logger.WithContext(ctx).Info("Request", zap.String("request_url", ctx.Request.URL.String()),
			zap.String("request_method", ctx.Request.Method),
			zap.Any("request_headers", ctx.Request.Header),
			zap.String("request_params", body))
		ctx.Next()
	}
}

func ResponseLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		startTime := time.Now()
		ctx.Next()
		duration := time.Since(startTime).String()
		logger.WithContext(ctx).Info("Response", zap.Any("response_body", blw.body.String()), zap.Any("time", duration))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
