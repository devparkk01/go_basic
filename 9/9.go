package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64 {
		"soup" : 56.33 , 
		"naan" : 65 ,
		"paneer masala" : 265, 
		"ice cream" : 200 ,
	}

	// fmt.Println(menu) 
	fmt.Println(menu["soup"]) 
	menu["soup"] = 400 
	fmt.Println(menu) 

	menu["noodles"] = 230 
	fmt.Println(menu) 

	// looping through the map
	for key ,value := range menu {
		fmt.Println(key , " : " , value) 
	}

	fmt.Println(len(menu)) 

}