package logger

import (
	"reflect"
	"time"
)

const (
	LogTypeSYS = "SYS"
)

const separator = "|"

var (
	TypeSliceOfBytes = reflect.TypeOf([]byte(nil))
	TypeTime         = reflect.TypeOf(time.Time{})
)
