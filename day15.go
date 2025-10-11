package main

import ("fmt";"strings")
func day15(){
	var slice = []string{"apple","apple", "banana", "apple", "orange", "banana", "grape", "apple"};
	var result = []string {};

	fmt.Println("Before removing dup the result is : " ,result);
		result = append(result, slice[0]);
		for x:= 1 ; x < len(slice) ; x++ {
			var isFound bool = false;


			for y:=0 ; y < len(result) ; y++ {
				fmt.Println("Im comparing ", slice[x], " to ", result[y] , "and its" , slice[x] != result[y]);
                 if strings.EqualFold(slice[x], result[y]) {
					isFound = true
					break;
				 }
				
			}
			if !isFound{
				result = append(result, slice[x]);
			 }
			 
			 fmt.Println("after that the result is now", result);

		}

fmt.Println("After removing all the dup result is : " ,result);
}	