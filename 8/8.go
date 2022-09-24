package main

import (
	"fmt"
	"math"
)
/* will look into functions here 

*/
func sayGreeting (name string ) (string) {
	var greeting  = "Hey " + name 
	return greeting  
}

func morning()(string) {
	return "very good morning to you ! "
}

func night()(string ) {
	return "Good night to you " 
}

//  making a callback function 
func greeting (name string ,day func() string ) (string) {
	var res string = "Hey " + name  + " , " 
	res += day() 
	return res 
}

// multiple return values

func circle (radius float64)(area float64 , circumference float64) {
	area =  math.Pi * radius * radius 
	circumference = math.Pi * 2 * radius 
	return 
}

func eligibility (age int)() {
	if age < 18 {
		fmt.Println("You are not eligible !")
	}else if age <= 60  {
		fmt.Println("You are eligible !")
	}else {
		fmt.Println("You are not eligible !")
	}
	
}

//  function to check whether the given string is palindrome or not 
func isPalindrome(s string) (bool) {
	if s == ""  {
		return true 
	} 
	i , j := 0 , len(s) - 1 

	for i < j  {
		if s[i] != s[j] {
			return false 
		} else {
			i++ 
			j--
		}
	}
	return true 

}




func main () {
	// var name = "Dev"
	// fmt.Println(sayGreeting(name)) 

	// fmt.Println(greeting(name , morning)) 
	// fmt.Println(greeting(name , night)) 

	// area , circumference := circle(10.00)
	// fmt.Println(area , circumference) 

	eligibility(64) 
	fmt.Printf("Palindrome %v : %v \n" , "Japan" , isPalindrome("Japan") )
	fmt.Printf("Palindrome %v : %v \n" , "" , isPalindrome("") )
	fmt.Printf("Palindrome %v : %v \n" , "LOL" , isPalindrome("LOL") )
	fmt.Printf("Palindrome %v : %v \n" , "interetni" , isPalindrome("interretni") )





}