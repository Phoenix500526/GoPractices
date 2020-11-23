// hello.go
package main

// #cgo LDFLAGS: -L. -lhello
// void SayHello(const char* s);
import "C"

func main(){
	C.SayHello(C.CString("Hello, World\n"))
}