package main

import ("fmt")
func day6(){

     var number int
	 var result int

    fmt.Print("Enter a Number: ")
	fmt.Scanln(&number)
    fmt.Print("The sum of digit", number)
	   for number >0{
             result += number%10
			 number/= 10
			}
	   
	fmt.Println(" is",result)
	
	






}