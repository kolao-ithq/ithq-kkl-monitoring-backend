package logs

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.StacktraceKey = ""

	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	Log = logger
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func GinZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Capture request body
		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		// Capture response body
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		c.Next()

		latency := time.Since(start)

		Log.Info("HTTP REQUEST",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Referer()),
			zap.Duration("latency", latency),
			zap.String("request_body", string(reqBody)),
			zap.String("response_body", blw.body.String()),
			zap.Int("bytes_sent", c.Writer.Size()),
			zap.String("host", c.Request.Host),
			zap.String("protocol", c.Request.Proto),
			zap.Strings("errors", c.Errors.Errors()),
		)
	}
}


func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg interface{}, fields ...zap.Field) {
	switch v := msg.(type) {
	case error:
		Log.Error(v.Error(), fields...)
	case string:
		Log.Error(v, fields...)
	}
}
