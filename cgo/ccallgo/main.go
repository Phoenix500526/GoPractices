package main
//#include "hello.h"
import "C"
import "fmt"

//export SayHello
func SayHello(s *C.char){
	fmt.Print("Go Say Hello:" + C.GoString(s))
}

func main(){
	C.SayHello(C.CString("Hello, World\n"))
}