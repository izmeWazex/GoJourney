package main

import ("fmt")


func day4(){

	var begining int = 1
	var end int = 100

	for x := begining ; x<= end ; x++ {

	if(((x % 3) == 0) && ((x% 5) == 0)){
		fmt.Println("FizzBuzz")
	}else if((x%3) == 0){
       fmt.Println("Fizz")
	}else if((x % 5) == 0){
		fmt.Println("Buzz")
	 }else{
		fmt.Println(x)

	 }
	}



}