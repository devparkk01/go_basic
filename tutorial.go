package main 
import "fmt" 


func add(x int , y int  ) (int) {
	return x + y 
}

func main() {
	fmt.Println("Hello Dev") 
	fmt.Println("Tum gande ho")
	res := add(10 , 13) 
	fmt.Println(res) 
}