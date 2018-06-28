// START OMIT
//     var a string
//     var done bool
// ----------------------
//     T1      |    T2
// ----------------------
// a = "hey"   | for !done {}
// done = true | print(a)
// END OMIT
package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
