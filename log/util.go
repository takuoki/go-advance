package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/fcfcqloow/go-advance/check"
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

func getContextValue(ctx context.Context, key interface{}) string {
	value := ctx.Value(key)
	if check.IsNil(value) {
		return fmt.Sprintf("%s: ", cnv.MustStr(key))
	}
	return fmt.Sprintf("%s: %s", cnv.MustStr(key), cnv.MustStr(value))

}
func getContextValues(ctx context.Context, keys []interface{}) (result string) {
	for _, key := range keys {
		result += getContextValue(ctx, key) + ", "
	}
	return
}
