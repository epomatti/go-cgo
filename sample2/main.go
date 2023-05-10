package main

import "fmt"

/*
	#include <stdio.h>
	#include <math.h>
	#include "calc.h"
*/
import "C"

func main() {
	fmt.Println(C.add(1, 2))
}
