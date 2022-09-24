package main 
import (
	"fmt"
)
// will look into pointers in this 
func update(name *string) () {
	*name += "DMG"
}

func main() {
	a := 10
	b := &a 
	fmt.Println("b stores : " , b )
	fmt.Println("after derefrencing : " , *b)

	name := "DEV" 
	update(&name)
	fmt.Println(name) 


}