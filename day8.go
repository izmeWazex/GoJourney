package main

import (
	"fmt"
)
func day8(){
	

	var number int
	isPrime := true
	squareRoot := 0
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)


	if ( number < 2){
		fmt.Println("The number you enter is prime number ")
		return
	   }

	for x:=1 ; x<= number; x++ {
		if x*x == number {
			squareRoot = x
           break
		}
		if x*x > number{
			squareRoot = x-1
			break
		}
	}
    
  
	for x:=2 ; x<=squareRoot ; x++ {
     if( number % x == 0) {
        isPrime = false
		break
	 }
	}
   

	if isPrime {
		fmt.Println("The number you enter is prime number ")

	}else{
		fmt.Println("The number you enter is not prime number ")

	}


}