// arrays and slices

package main 
import ( "fmt" ) 


func main () {
	var ages[3]int = [3]int{20 , 18 , 19}  
	fmt.Println(ages) 

	// empty array of size 3 
	var names = [3]string{}
	fmt.Println(names) 

	names[0] = "Dev"
	names[1] = "Harsh"
	names[2] = "Ed"
	fmt.Println(names) 

	// can't do this 
	// names[3] = "Rahul" 
	// fmt.Println(names)

	articles := [3]float64{23 , 30 , 50} 
	fmt.Println(articles) 

	fmt.Println("length : " , len(names) , " , capacity :" , cap(names)) 

	// slices

	var marks[]int = []int{} ; 
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  

	marks = append(marks , 10)
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  

	marks = append(marks , 20 ) 
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  

	marks = append(marks , 30) 
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  


	marks = append(marks , 40) 
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  

	marks = append(marks , 50) 
	fmt.Println("length : " , len(marks) , " , capacity :" , cap(marks))  






}