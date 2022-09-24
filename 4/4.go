package main 
import ("fmt")

/*
more on arrays and slices 
*/

func main() {
	// initializing an array 
	authors := [8]string{"hessan" , "mike"  , "Gabour" , "Sheru" , "hemawari" , "Shwaki" , "Chims"} 
	fmt.Println(authors)
	
	// initializing a slice from the above array 
	topAuthors := authors[1:3] 
	fmt.Println(topAuthors) 

	topAuthors[0] = "Mike"
	fmt.Println(topAuthors)
	fmt.Println(authors) 

	bestAuthors := authors[:5] 
	fmt.Println(bestAuthors)

	bestAuthors[0] = "Hessan"
	fmt.Println(bestAuthors)
	fmt.Println(authors)

	bestAuthors[1] = "Mikey"
	fmt.Println(bestAuthors)
	fmt.Println(topAuthors)
	fmt.Println(authors)



}