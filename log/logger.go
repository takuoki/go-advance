package log

import (
	"errors"
	"log"
	"strings"

	cnv "github.com/fcfcqloow/go-advance/convert"
)

type (
	LogLevel    int
	FilePath    string
	ProgramLine int
	MethodName  string
	OutputFunc  func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string
)

const (
	LOG_LEVEL_TRACE     LogLevel = 401
	LOG_LEVEL_DEBUG     LogLevel = 400
	LOG_LEVEL_INFO      LogLevel = 300
	LOG_LEVEL_WARN      LogLevel = 200
	LOG_LEVEL_ERROR     LogLevel = 100
	LOG_LEVEL_EMERGENCY LogLevel = -1
	LOG_LEVEL_FATAL     LogLevel = -2
	LOG_LEVEL_DISRUPT   LogLevel = -999
	DEFAULT_SKIP        int      = 2
	DEFAULT_LOG_LEVEL   LogLevel = LOG_LEVEL_INFO
)

var (
	logLevel    = DEFAULT_LOG_LEVEL
	skip        = DEFAULT_SKIP
	keys        = []interface{}{}
	SetWriter   = log.SetOutput
	TraceOutput = func(values []interface{}, filePath FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(filePath), " / ", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	DebugOutput = func(values []interface{}, filePath FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(filePath), " / ", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	InfoOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	WarnOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	ErrorOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	EmergencyOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	FatalOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
	DisruptOutput = func(values []interface{}, _ FilePath, line ProgramLine, name MethodName) []string {
		return []string{"[", string(name), "]: ", strings.Join(InterfacesToStrings(values), " ")}
	}
)

func ParseLevel(level string) (LogLevel, error) {
	switch strings.ToLower(level) {
	case "trace", "t", "tra":
		return LOG_LEVEL_TRACE, nil
	case "debug", "deb", "d", "develop", "dev":
		return LOG_LEVEL_DEBUG, nil
	case "info", "inf", "i":
		return LOG_LEVEL_INFO, nil
	case "warn", "waring", "w":
		return LOG_LEVEL_WARN, nil
	case "error", "err", "e":
		return LOG_LEVEL_ERROR, nil
	case "emergency":
		return LOG_LEVEL_EMERGENCY, nil
	case "fatal", "f":
		return LOG_LEVEL_FATAL, nil
	case "disrupt":
		return LOG_LEVEL_DISRUPT, nil
	default:
		return LogLevel(-1000), errors.New("No match pattern")
	}
}
func SetFunctionSkip(_skip int) {
	skip = _skip
}
func SetLevel(level LogLevel) {
	logLevel = level
}
func SetLevelOrDefault(level interface{}, defaultLevel LogLevel) {
	if _level, err := ParseLevel(cnv.MustStr(level)); err != nil {
		SetLevel(defaultLevel)
	} else {
		SetLevel(_level)
	}
}

func GetLevel() LogLevel {
	return logLevel
}

func Trace(values ...interface{}) {
	println(LOG_LEVEL_TRACE, log.Ldate|log.Ltime|log.Lmicroseconds, "[TRACE] ", TraceOutput, values...)
}
func Debug(values ...interface{}) {
	println(LOG_LEVEL_DEBUG, log.Ldate|log.Ltime, "[DEBUG] ", DebugOutput, values...)
}
func Info(values ...interface{}) {
	println(LOG_LEVEL_INFO, log.Ldate|log.Ltime, "[INFO] ", InfoOutput, values...)
}
func Warn(values ...interface{}) {
	println(LOG_LEVEL_WARN, log.Ldate|log.Ltime, "[\u001B[33mWARN\u001B[0m] ", WarnOutput, values...)
}
func Error(values ...interface{}) {
	println(LOG_LEVEL_WARN, log.Ldate|log.Ltime|log.Lmicroseconds, "[\u001B[31mERROR\u001B[0m] ", ErrorOutput, values...)
}
func Emergency(values ...interface{}) {
	fatalln(LOG_LEVEL_EMERGENCY, log.Ldate|log.Ltime|log.Lmicroseconds, "[EMERGENCY] ", EmergencyOutput, values...)
}
func Disrupt(values ...interface{}) {
	fatalln(LOG_LEVEL_DISRUPT, log.Ldate|log.Ltime|log.Lmicroseconds, "[DISRUPT] ", DisruptOutput, values...)
}
func Fatal(values ...interface{}) {
	println(LOG_LEVEL_FATAL, log.Ldate|log.Ltime|log.Lmicroseconds, "[FATAL] ", FatalOutput, values...)
}
