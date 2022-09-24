package main

import (
	"fmt"
	"strconv"
)


func main () {
	var age string = "34"
	// ParseInt function
	val, err  := strconv.ParseInt(age , 0 , 64 )
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v : %T\n" , val , val )
	}
	// ParseFloat function 
	
	amount , err := strconv.ParseFloat("-353.23" , 64) 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Parsed float is : " , amount) 
	}

	// Atoi function
	res , _ := strconv.Atoi("-200")
	fmt.Println("Parsed int value : " , res ) 

	// Itoa function 
	
	str  := strconv.Itoa(34)
	fmt.Printf("Converted string is : %v and its type : %T\n" , str , str )


}