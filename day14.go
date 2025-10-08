package main

import ("sort";
         "fmt";
		"strings")

func sortWord(word string) string {
    var result []string;
	for x:=0 ; x<len(word); x++ {
		result = append(result, strings.ToLower(string(word[x])));
	}


    sort.Strings(result)
	return strings.Join(result,"");

}

func day14(){

	var firstWord string = "lisTen";
	var secondWord string = "silent";

	fmt.Print("The two word is ")
     if len(firstWord) != len(secondWord){
		fmt.Println("not a anagram")
	 }else if sortWord(firstWord) == sortWord(secondWord){
		fmt.Println("anagram");
	}else{
		fmt.Println("not a anagram");
	}

      

}