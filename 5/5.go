package main 
import (
	"fmt"
	"strings"
)
/*
will look into standard library 'strings'
*/

func main() {
	var text = "max heap by default in priority queue" 

	fmt.Println(text ,  "Length : " , len(text )) 


	// Compare function : Compare returns an integer comparing two strings lexicographically. 
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b. 

	var res int ; // declaring an int  variable 
	res = strings.Compare("DEV" , "RAHUL") 
	fmt.Println(res) 

	res = strings.Compare("RAHUL" , "Harsh") 
	fmt.Println(res) 

	res =  strings.Compare("DEV" , "DEV") 
	fmt.Println(res)


	// Contains function : returns bool 
	var ans bool // declaring a bool variable 

	ans = strings.Contains(text, "max") 
	fmt.Println(ans) 

	ans = strings.Contains(text, "maximum") 
	fmt.Println(ans) 

	// Count function :
	res = strings.Count(text , "ue")
	fmt.Println("Occurrence of ue : " , res) 

	res = strings.Count(text , "e" )
	fmt.Println("Occurrence of e : " , res) 
	

	// Index function : returns the index of the first instance of 
	//  substr in s, or -1 if substr is not present in s. 

	res = strings.Index(text , "ea") 
	fmt.Println("Index of first instance of ea : " , res) 

	res = strings.Index(text , "power" ) 
	fmt.Println("Index of first instance of power : " , res) 

	// Split function , returns a slice 

	sliced := strings.Split(text , " ") 
	fmt.Println(sliced , " , Length : " , len(sliced)) 

	sliced = strings.Split(text , "") 
	fmt.Println(sliced , " , Length : " , len(sliced ))

	// Join function , returns a String after joining the array/slice with separator after each element

	newSlice := []string{"Dev" , "Harsh" , "Ram" , "Ed"}

	newJoinedString := strings.Join(newSlice , " ")
	fmt.Println(newJoinedString) 

	newJoinedString = strings.Join(newSlice , "##")
	fmt.Println(newJoinedString) 

	// Title function 
	fmt.Println(strings.Title(text)) 

	// ToUpper function 
	fmt.Println(strings.ToUpper(text))

	// ToLower function
	fmt.Println(strings.ToLower(text)) 
	
	// TrimSpace function 
	fmt.Println(strings.TrimSpace("   Dev     "))


}