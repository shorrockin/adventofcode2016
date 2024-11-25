package assert

import (
	"fmt"
	"os"
	"runtime/debug"
)

func logAssert(msg string, data ...any) {
	fmt.Fprintf(os.Stderr, "ASSERT:")
	fmt.Fprintf(os.Stderr, "\tmsg=%s\n", msg)

	for i := 0; i < len(data); i += 2 {
		fmt.Fprintf(os.Stderr, "\t%v=%+v\n", data[i], data[i+1])
	}
	fmt.Fprintf(os.Stderr, string(debug.Stack()))
	os.Exit(1)
}

func Fail(msg string, data ...any) {
	logAssert(msg, data...)
}

func Assert(truth bool, msg string, data ...any) {
	if truth {
		return
	}
	logAssert(msg, data...)
}

func Refute(falsehood bool, msg string, data ...any) {
	if !falsehood {
		return
	}
	logAssert(msg, data...)
}

func NotNil(value any, msg string, data ...any) {
	if value != nil {
		return
	}
	logAssert(msg, data...)
}

func Nil(value any, msg string, data ...any) {
	if value == nil {
		return
	}
	logAssert(msg, data...)
}

func NoError(err error, msg string, data ...any) {
	if err == nil {
		return
	}
	data = append(data, err)
	logAssert(msg, data...)
}

func False(value bool, msg string, data ...any) {
	if !value {
		return
	}
	logAssert(msg, data...)
}

func True(value bool, msg string, data ...any) {
	if value {
		return
	}
	logAssert(msg, data...)
}

func Equal(expected, actual any, msg string, data ...any) {
	if expected == actual {
		return
	}
	data = append(data, "expected", expected, "actual", actual)
	logAssert(msg, data...)
}

func NotEqual(expected, actual any, msg string, data ...any) {
	if expected != actual {
		return
	}
	data = append(data, "expected", expected, "actual", actual)
	logAssert(msg, data...)
}
