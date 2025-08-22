package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zsljava/gokit/common/response"
	"github.com/zsljava/gokit/util/jwt"
	"github.com/zsljava/gokit/util/log"
	"go.uber.org/zap"
	"net/http"
)

var ErrUnauthorized = response.NewError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))

func StrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			logger.WithContext(ctx).Warn("No token", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			response.Error(ctx, ErrUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("token common", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}), zap.Error(err))
			response.Error(ctx, ErrUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func NoStrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	if userInfo, ok := ctx.MustGet("claims").(*jwt.MyCustomClaims); ok {
		logger.WithValue(ctx, zap.String("user_id", userInfo.UserId))
	}
}
