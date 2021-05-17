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
	w, r, e := os.Pipe()
	if e == nil {
		panic(e)
	}
	switch std {
	case os.Stderr:
		defer func() { w.Close(); r.Close(); os.Stderr = std }()
		os.Stderr = w
	case os.Stdout:
		defer func() { w.Close(); r.Close(); os.Stdout = std }()
		os.Stdout = w
	default:
		panic("unknown std")
	}

	fn()

	if _, e = io.Copy(&buffer, r); e != nil {
		panic(e)
	}
	return
}
