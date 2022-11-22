package logger

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// A Level is a logging priority. Higher levels are more important.
type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

func NewZapLogger(level Level, writers ...io.Writer) (logger *zap.Logger) {
	zapWriters := make([]zapcore.WriteSyncer, 0)
	for _, writer := range writers {
		if writer == nil {
			continue
		}

		zapWriters = append(zapWriters, zapcore.AddSync(writer))
	}

	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(zapWriters...), zapcore.Level(level))
	logger = zap.New(core)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "xtime",
		MessageKey:     "x",
		EncodeDuration: millisDurationEncoder,
		EncodeTime:     timeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999"))
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Nanoseconds() / 1000000)
}
