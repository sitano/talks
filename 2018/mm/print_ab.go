// START OMIT
//      var a, b int
// ----------------------
//     T1    |      T2
// ----------------------
//   a = 1   |  print(b)
//   b = 2   |  print(a)
// END OMIT
package main

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
