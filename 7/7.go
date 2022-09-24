package main 
import ("fmt")

// will look into loops in this code

func main() {
	// kinda while loop
	i := 0 
	for i < 4 {
		fmt.Println(i) 
		i++ 
	}

	// proper  for loop
	for i :=0 ; i < 4 ; i++ {
		fmt.Println(i)
	}

	names := []string{"Dev" , "Rahul" , "Harsh"}
	names = append(names , "Ed")

	for index , value := range names {
		fmt.Println(index , value) 
	}

	for _ , value := range names {
		fmt.Println(value) 
	}

	


}