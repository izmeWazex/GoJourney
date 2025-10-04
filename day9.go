package main


import ("fmt")
func day9(){

	
	number := []int  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	left:= 0
	right:= len(number) - 1

	for {
		if left >= right {
			break
		}
		// I make more challenging no using temporary variable
		// temp:= number[left]
		// number[left] = number[right] + number[left]
		// number[right] = temp
		number[left] = number[left] + number[right]
		number[right] = number[left] - number[right]
		number[left] = number[left] - number[right]
		
		left++
		right--
	}
    fmt.Println(number)



	
}