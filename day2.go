package main

import ("fmt")

func day2(){
	var int1 int = 5
	var int2 int = 3
	fmt.Println("BEFORE\nINTIGER 1:", int1 , "\nInteger 2:", int2)
    int1 = int1 + int2
	int2 = int1 - int2
	int1 = int1 - int2

    fmt.Println("AFTER\nINTIGER 1:", int1 , "\nInteger 2:", int2)

}