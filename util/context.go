package util

import (
	"context"
	"github.com/zsljava/gokit/constant"
)

func GetTraceId(ctx context.Context) string {
	return ctx.Value(constant.TraceIDKey).(string)
}
