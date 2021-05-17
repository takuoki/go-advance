package xos

import (
	"bytes"
	"io"
	"os"
)

func CaptureStderr(fn func()) (buffer bytes.Buffer) {
	return capture(os.Stderr, fn)
}

func CaptureStdout(fn func()) (buffer bytes.Buffer) {
	return capture(os.Stdout, fn)
}

func capture(std *os.File, fn func()) (buffer bytes.Buffer) {
	r, w, e := os.Pipe()
	if e != nil {
		panic(e)
	}
	switch std {
	case os.Stderr:
		defer func() { os.Stderr = std }()
		os.Stderr = w
	case os.Stdout:
		defer func() { os.Stdout = std }()
		os.Stdout = w
	default:
		panic("unknown std")
	}

	fn()
	w.Close()
	if _, e = io.Copy(&buffer, r); e != nil {
		panic(e)
	}
	return
}
