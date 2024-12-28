package logger

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	start       time.Time
	last        time.Time
	name        string
	laps        int
	indentation int
}

type LogOptions struct {
	includeTotal bool
	includeDelta bool
	newline      bool
	indent       bool
	variables    map[string]interface{}
}

type Option func(*Logger, *LogOptions)

var names = make(map[string]int)

func New(name string) Logger {
	names[name]++
	if names[name] != 1 {
		name = fmt.Sprintf("%s-%d", name, names[name])
	}
	start := time.Now()
	bm := Logger{start, start, name, 0, 0}
	bm.Log("Starting", ExcludeDelta)
	return bm
}

func (logger *Logger) Reset(msg string, options ...interface{}) {
	logger.start = time.Now()
	logger.last = logger.start
	logger.laps = 0
	logger.Log(msg, append(options, ExcludeDelta)...)
}

func (logger *Logger) Checkpoint(msg string, options ...interface{}) {
	logger.Log(msg, options...)
	logger.last = time.Now()
}

func LogReturn[T any](logger *Logger, value T, options ...interface{}) T {
	logger.Log("returning", append(options, With("value", value), IncludeTotal)...)
	return value
}

func (logger *Logger) Log(msg string, opts ...interface{}) {
	options := []Option{}
	for i := 0; i < len(opts); i++ {
		if i+1 < len(opts) {
			if name, ok := opts[i].(string); ok {
				if value, ok := opts[i+1].(interface{}); ok {
					options = append(options, With(name, value))
					i++
					continue
				}
			}
		}
		if option, ok := opts[i].(Option); ok {
			options = append(options, option)
		}
	}

	logOptions := &LogOptions{
		includeTotal: false,
		includeDelta: true,
		indent:       false,
		newline:      false,
		variables:    make(map[string]interface{}),
	}
	for _, option := range options {
		option(logger, logOptions)
	}
	elapsed := time.Since(logger.last)
	total := time.Since(logger.start)
	logger.laps++

	indentation := ""
	for range logger.indentation {
		indentation += "  "
	}
	if logOptions.indent {
		indentation += "  "
	}

	trailing := []string{}
	for name, value := range logOptions.variables {
		trailing = append(trailing, fmt.Sprintf("%s=%+v", name, value))
	}

	if logOptions.includeDelta {
		trailing = append(trailing, fmt.Sprintf("Δ=%s", duration(elapsed)))
	}
	if logOptions.includeTotal {
		trailing = append(trailing, fmt.Sprintf("∑=%s", duration(total)))
	}

	newline := ""
	if logOptions.newline {
		newline = "\n"
	}

	fmt.Printf("%v%d. [%s] %s%s %s\n", newline, logger.laps, logger.name, indentation, msg, strings.Join(trailing, ", "))
}

func duration(d time.Duration) string {
	if d < time.Microsecond {
		// Display in nanoseconds
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Millisecond {
		// Display in microseconds with one decimal
		return fmt.Sprintf("%.1f µs", float64(d.Nanoseconds())/1000.0)
	} else if d < time.Second {
		// Display in milliseconds with one decimal
		return fmt.Sprintf("%.1f ms", float64(d.Milliseconds()))
	} else {
		// Display in seconds with two decimals
		return fmt.Sprintf("%.2f s", d.Seconds())
	}
}

func IndentOnce(logger *Logger, options *LogOptions) {
	options.indent = true
}

func Indent(logger *Logger, options *LogOptions) {
	logger.indentation++
}

func Unindent(logger *Logger, options *LogOptions) {
	logger.indentation--
}

func IncludeTotal(logger *Logger, options *LogOptions) {
	options.includeTotal = true
}

func ExcludeDelta(logger *Logger, options *LogOptions) {
	options.includeDelta = false
}

func WithNewline(logger *Logger, options *LogOptions) {
	options.newline = true
}

func With(name string, value interface{}) Option {
	return func(logger *Logger, options *LogOptions) {
		options.variables[name] = value
	}
}
