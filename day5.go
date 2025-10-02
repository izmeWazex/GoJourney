package main
import ("fmt")
func day5(){
	var number int

 fmt.Print("Enter a number for factoring: ")
 fmt.Scanln(&number)

   result := calculateFactorial(number)
   fmt.Println("The factorial of ", number , "is ", result)
  
}

func calculateFactorial(number int) []int{
    result := make([]int, number)
    for x:= 0;x<=number-1; x++{
		result[x] = helper(x+1)
	}
	return result

}
func helper(number int) int{
 if number ==1 {
	return 1
 }else{
	return number * helper(number-1)
 }

}