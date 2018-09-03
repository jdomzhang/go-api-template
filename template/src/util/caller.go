package util

import "runtime"

func GetCurrentMethodName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	// fmt.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
	return frame.Function
}
