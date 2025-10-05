package main

import ("fmt")

func day10(){
    slice := []int {2,5,7,4,8,23,67,5,87,56};
	sliceLen := len(slice);
	highest := 0;

	for x:=0 ; x<=sliceLen-1; x++ {
		if highest < slice[x] {
			highest = slice[x];
		}
	}
	fmt.Println("The highest is ", highest)

}