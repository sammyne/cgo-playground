package hello

/*
extern int helloFromC();
*/
import "C"

func World() {
	//call c function
	C.helloFromC()
}
