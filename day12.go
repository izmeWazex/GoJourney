package main

import ("fmt";
        "strings")

func isVowel (ch string) bool {
	switch strings.ToLower(ch) {
	case "a", "e", "i", "o", "u":
     return  true;
	default:
		return false;
	}
}

func day12(){

    var vowel int = 0;
	var word string;
	fmt.Println("enter a word");
	fmt.Scanln(&word);

	for x:=0; x<= len(word)-1; x++ {
       if isVowel(string(word[x])) {
         vowel++;
	   }
	}


	fmt.Println("The total number of vowel is", vowel);

   
	



}