package main

import (
	"fmt"
)

func day7(){

	var word string;

	fmt.Print("Enter word to check palyndrome: ")
	fmt.Scanln(&word)
	
	left := 0
	right := len(word) -1
	var result bool = true;

		for left <= right{
			if word[left] != word[right] {
				result = false;
				break;
			}
			left++
			right--
		} 
	fmt.Print("The word ", word, " is a ")

		if result {
			fmt.Println("palyndrome")
		}else{
			fmt.Println("not palyndrome")
		}

}