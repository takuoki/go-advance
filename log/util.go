package log

import (
	"log"
	"path/filepath"
	"runtime"
	"strconv"

	cnv "github.com/fcfcqloow/go-advance/convert"
)

func InterfacesToStrings(arg []interface{}) []string {
	stringArray := make([]string, len(arg))
	for i, v := range arg {
		stringArray[i] = cnv.MustStr(v)
	}
	return stringArray
}

func getInfo() (FilePath, ProgramLine, MethodName) {
	pt, file, line, _ := runtime.Caller(skip)
	return FilePath(file + ":" + strconv.Itoa(line)), ProgramLine(line), MethodName(filepath.Base(runtime.FuncForPC(pt).Name()))
}

func println(level LogLevel, flags int, prefix string, output OutputFunc, values ...interface{}) {
	if should(level) {
		filePath, line, name := getInfo()
		log.SetFlags(flags)
		log.SetPrefix(prefix)
		log.Println(output(values, filePath, line, name))
	}
}
func fatalln(level LogLevel, flags int, prefix string, output OutputFunc, values ...interface{}) {
	if should(level) {
		filePath, line, name := getInfo()
		log.SetFlags(flags)
		log.Fatalln("\u001B[31m", prefix, output(values, filePath, line, name), "\x1b[0m")
	}
}
func should(level2 LogLevel) bool {
	return logLevel >= level2
}
