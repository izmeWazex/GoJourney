package main
import ("fmt")

func day3(){

	var number int = 4
	
    fmt.Print("The number ", number, " is ")
	if ((number % 2 ) == 0){
	fmt.Println("even")
	}else{
	fmt.Println("odd")
	}
}