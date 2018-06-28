// START OMIT
//      var a, b int
// ----------------------
//     T1    |      T2
// ----------------------
//   a = 1   |  print(b)
//   b = 2   |  print(a)
// END OMIT
package main

import (
	"runtime"
	"runtime/debug"
)

type mem struct {
	a, b int
}

func (m *mem) t1() {
	m.a = 1
	m.b = 2
}

func (m *mem) t2() {
	s := []byte("00\n")
	s[0] = byte('0' + m.b)
	s[1] = byte('0' + m.a)
	print(string(s))
}

func main() {
	runtime.GOMAXPROCS(5)
	debug.SetGCPercent(-1)
	for i := 0; i < 1000000; i++ {
		m := new(mem)
		go m.t1()
		go m.t2()
	}
}
