package driver

import "fmt"

type SafeGoroutineTester struct{}

func (s *SafeGoroutineTester) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}
func (s *SafeGoroutineTester) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	panic("TEST FATAL FAILURE")
}
