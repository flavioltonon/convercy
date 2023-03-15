package logging

import (
	"fmt"
	"strconv"
)

type Logger interface {
	Info(message string, fields ...Field)
	Error(message string, fields ...Field)
	Fatal(message string, fields ...Field)
}

type Field struct {
	key   string
	value string
}

func String(key string, value string) Field {
	return Field{
		key:   key,
		value: value,
	}
}

func Int(key string, value int) Field {
	return String(key, strconv.Itoa(value))
}

func Stringer(key string, value fmt.Stringer) Field {
	return String(key, value.String())
}

func Error(value error) Field {
	return String("error", value.Error())
}

func (f Field) Key() string {
	return f.key
}

func (f Field) Value() string {
	return f.value
}
