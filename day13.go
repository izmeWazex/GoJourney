package main

import (
	"fmt"
	"math"
)

func day13(){


	var number int;
	var numberSummation int;
	fmt.Println("Enter a digit number");
	fmt.Scanln(&number);
	
	var numberOfDigit int = int(len(fmt.Sprint(number)));
	var storedNumber int = number;

	for number > 0 {
		digitInNum :=number % 10;
		numberSummation += int(math.Pow(float64(digitInNum), float64(numberOfDigit)));
		number /= 10;
	}
    
    fmt.Print("The number ",storedNumber," is ");
	if storedNumber == numberSummation {
      fmt.Println("Armstrong Number");
	}else {
		fmt.Println("not a Armstrong Number");
	}
	



}