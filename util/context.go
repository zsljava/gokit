package util

import (
	"context"
	"github.com/zsljava/gokit/global"
)

func GetTraceId(ctx context.Context) string {
	return ctx.Value(global.TraceIDKey).(string)
}
