package main

import (
	"fmt"
)

func day11(){

	var word string ;
	fmt.Println("Enter a word: ");
	fmt.Scanln(&word);

	for x:=0; x<=len(word)-1; x++ {
		for y:=x+1; y<=len(word)-2; y++ {
          if string(word[x] ) == string(word[y] ) {
			fmt.Println("The word",word,"is not all unique character" );
			return;
		  }
		}
	}
	fmt.Println("The word",word,"is all unique character" );


}