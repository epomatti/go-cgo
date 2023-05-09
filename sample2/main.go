package main

import "fmt"

/*
	#include <stdio.h>
	#include <math.h>
	#include "mylib.h"
*/
import "C"

func main() {
	fmt.Println(C.multiply(10, 20))
}
