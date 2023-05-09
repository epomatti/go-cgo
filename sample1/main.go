package main

/*
#include <stdio.h>
#include <stdlib.h>

 const char* hello_world() {
     return "Hello World!";
 }
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println(C.GoString(C.hello_world()))
}
