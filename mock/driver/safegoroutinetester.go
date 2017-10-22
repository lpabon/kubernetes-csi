/*
Copyright 2017 Luis Pabón lpabon@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
