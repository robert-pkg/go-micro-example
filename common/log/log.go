package log

import (
	"time"

	zapLog "common/log/zap-wrapper"

	"github.com/micro/go-micro/v2/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(fileName string) {
	zCfg := zap.NewProductionConfig()

	zCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	zCfg.Encoding = "json"
	zCfg.EncoderConfig.TimeKey = "t"
	zCfg.EncoderConfig.LevelKey = "l"
	zCfg.EncoderConfig.NameKey = "logger"
	zCfg.EncoderConfig.CallerKey = "c"
	zCfg.EncoderConfig.MessageKey = "msg"
	zCfg.EncoderConfig.StacktraceKey = "st"
	zCfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
	zCfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	zCfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	zCfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder //zapcore.ShortCallerEncoder

	log, err := zapLog.NewLogger(
		zapLog.WithCallerSkip(2),
		zapLog.WithConfig(zCfg),
		zapLog.WithLogFileName(fileName),
	)
	if err != nil {
		panic(err)
	}
	logger.DefaultLogger = log
}
