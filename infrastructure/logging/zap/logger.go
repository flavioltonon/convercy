package zap

import (
	"convercy/shared/logging"

	"go.uber.org/zap"
)

type Logger struct {
	core *zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &Logger{core: logger}, nil
}

func (l *Logger) Info(message string, fields ...logging.Field) {
	l.core.Info(message, l.newFields(fields)...)
}

func (l *Logger) Error(message string, fields ...logging.Field) {
	l.core.Error(message, l.newFields(fields)...)
}

func (l *Logger) Fatal(message string, fields ...logging.Field) {
	l.core.Fatal(message, l.newFields(fields)...)
}

func (l *Logger) newField(input logging.Field) zap.Field {
	return zap.String(input.Key(), input.Value())
}

func (l *Logger) newFields(inputs []logging.Field) []zap.Field {
	outputs := make([]zap.Field, 0, len(inputs))

	for _, input := range inputs {
		outputs = append(outputs, l.newField(input))
	}

	return outputs
}
